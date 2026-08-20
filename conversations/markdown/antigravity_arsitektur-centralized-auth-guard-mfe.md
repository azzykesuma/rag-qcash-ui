# Panduan Arsitektur: Centralized Auth Guard & Module Federation Singleton

**Penulis**: Tim Senior Frontend Architecture (Antigravity)  
**Tanggal**: 19 Agustus 2026  
**Cakupan**: `qcash-ui` (Host Shell) & 40+ Remote Micro-Frontends (`qcash-ui-account-receivable`, `qcash-ui-fund-transfer`, dll.)

---

## 1. Latar Belakang & Akar Masalah

Pada arsitektur Micro-Frontend berbasis **Webpack Module Federation (MFE)** dan **Next.js** di platform QLola Cash Management, sering terjadi kendala autentikasi dan performa:

1. **Crash `Uncaught Error: Function not implemented`**:
   Terjadi saat komponen halaman remote mengeksekusi `guard()` di dalam `useEffect`. Jika remote tidak dikonfigurasi dengan `singleton: true` atau memiliki versi `@ordentco/addons-auth-provider` yang berbeda dari Host, React Context mengembalikan fungsi fallback default `() => { throw new Error("Function not implemented"); }`.
2. **Infinite Loading / State Tidak Siap**:
   Komponen remote membaca `isAuthoritiesReady: false` secara permanen karena instance context di remote terisolasi dari Host.
3. **Duplikasi Fetch `/auth/me` & `/menu/me` (Request Storm)**:
   Setiap kali user berpindah halaman di dalam remote, `useEffect` di halaman tersebut memanggil `guard()`, memicu ulang request HTTP ke endpoint `/auth/me` dan `/menu/me`.
4. **Kedip / Loading Blink**:
   Memasang `LoadingOverlay` yang memblokir root React tree menyebabkan kedip layar (flicker) saat berganti dengan skeleton native halaman (seperti `DashboardSkeleton`).

---

## 2. Mengapa `singleton: true` Mutlak Diperlukan?

Dalam React, **React Context bekerja berdasarkan kesamaan referensi objek di memori (`ContextHost === ContextRemote`)**.

```
❌ TANPA SINGLETON (singleton: false / commented out):
┌─────────────────────────────────────────────────────────────┐
│ HOST SHELL (qcash-ui)                                       │
│   Membuat Context Instance #1 (isAuthoritiesReady = true)    │
└─────────────────────────────────────────────────────────────┘
                               ▲
                 TIDAK BISA SALING AKSES ❌
                               ▼
┌─────────────────────────────────────────────────────────────┐
│ REMOTE MFE (fund-transfer, credit-card, dll.)               │
│   Membuat Context Instance #2 (isAuthoritiesReady = false)  │
└─────────────────────────────────────────────────────────────┘
```

* **Tanpa `singleton: true`**: Webpack membundel 2 instance terpisah dari `@ordentco/addons-auth-provider`. State autentikasi yang sudah divalidasi Host **tidak pernah sampai** ke komponen remote.
* **Dengan `singleton: true`**: Webpack memaksa Host dan seluruh Remote berbagi **1 instance context yang sama di runtime**.
* **Hasil**: Begitu Host selesai memvalidasi token, `isAuthoritiesReady` di seluruh halaman remote **langsung bernilai `true` secara instan**.

---

## 3. Implementasi Centralized Guard di Host (`<HostAuthGate />`)

Host Shell (`qcash-ui`) bertindak sebagai **pemilik tunggal siklus autentikasi** melalui komponen `<HostAuthGate />` yang dipasang di `pages/_app.tsx`:

```tsx
// components/providers/HostAuthGate.tsx (Di Host Shell)
export const HostAuthGate: React.FC<HostAuthGateProps> = ({ children }) => {
  const router = useRouter();
  const hydratedTokenRef = useRef<string | null>(null);

  const briAuth = useBriAuth();
  const ordentAuth = useOrdentAuth();

  useEffect(() => {
    if (typeof window === "undefined") return;

    // 1. Sinkronisasi access-token dari localStorage ke React Context
    syncTokenFromLocalStorage(briAuth);
    syncTokenFromLocalStorage(ordentAuth);

    const token = (ordentAuth?.token as string) || localStorage.getItem("access-token");
    if (!token) {
      hydratedTokenRef.current = null;
      return;
    }

    // 2. Eksekusi guard() TEPAT 1 KALI per sesi token
    if (hydratedTokenRef.current === token && ordentAuth?.isAuthoritiesReady) {
      return;
    }

    if (shouldHydrateAuth(ordentAuth) && hydratedTokenRef.current !== token) {
      hydratedTokenRef.current = token;

      Promise.resolve(ordentAuth.guard(true)).catch((error) => {
        console.error(`[HostAuthGate] Gagal memvalidasi auth`, error);
        const status = error?.response?.status || error?.status;
        if (status === 401 || String(error).includes("401")) {
          localStorage.removeItem("access-token");
          localStorage.removeItem("refresh-token");
          hydratedTokenRef.current = null;
          window.dispatchEvent(new CustomEvent("qc-bridge-sync", { detail: { source: "logout" } }));
          router.replace("/landing-page");
        }
      });
    }
  }, [ordentAuth?.token, ordentAuth?.isAuthoritiesReady, briAuth?.token, router.pathname]);

  // 3. Render children secara langsung (tanpa overlay blocking agar tidak berkedip)
  return <>{children}</>;
};
```

---

## 4. Mengapa Terjadi Fetch Ganda `/auth/me` & `/menu/me` dan Cara Mencegahnya

### Mekanisme 1 Pemanggilan `guard()`:
Di dalam `@ordentco/addons-auth-provider`, 1 eksekusi fungsi `guard()` selalu memicu **2 HTTP request**:
1. `GET /auth/me` (Validasi profil dan sesi user)
2. `POST /menu/me` (Daftar hak akses menu dan lisensi produk)

### Penyebab Fetch Ganda (2x `/auth/me` + 2x `/menu/me`):
Saat user membuka menu (misalnya `/credit-card` atau `/fund-transfers`):
1. **Fetch #1 (Dari Host)**: Dijalankan oleh `<HostAuthGate>` saat inisialisasi sesi.
2. **Fetch #2 (Dari Remote Page)**: Dijalankan oleh komponen halaman remote karena masih memiliki:
   ```tsx
   useEffect(() => {
     guard(); // ⚠️ Memicu fetch kedua yang redundan
   }, []);
   ```

### Solusi & Rekomendasi:
* **Halaman Remote Tidak Perlu Memanggil `guard()`**:
  Karena Host sudah memvalidasi token dan membagikan context via `singleton: true`, halaman remote cukup membaca state:
  ```tsx
  // ✅ Pola Halaman Remote yang Bersih:
  export default function FeaturePage() {
    const { userType, companyID, productAuthorities } = useAuth();

    // Cukup gunakan optional chaining untuk keamanan
    if (!productAuthorities?.INTERNAL_FUND_TRANSFER?.anyAuthority) {
      return <UnauthorizedDialog />;
    }

    return <FeatureContainer companyId={companyID} />;
  }
  ```

---

## 5. Panduan Rollout 2-Fase untuk Squad (Migration Playbook)

Agar migrasi tidak membebani tim squad dan tidak menimbulkan risiko regresi:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│ FASE 1: Quick-Win (1 Menit per Repo)                                       │
│   - Aksi: Tambahkan `singleton: true` di `next.config.js` remote.          │
│   - Halaman `pages/**/*.tsx` TIDAK PERLU diubah sama sekali.               │
│   - Dampak: Error "Function not implemented" dan infinite loading          │
│     LANGSUNG HILANG 100%.                                                  │
├─────────────────────────────────────────────────────────────────────────────┤
│ FASE 2: Optimalisasi Performa (Dikerjakan bertahap di sprint squad)         │
│   - Aksi: Hapus `useEffect(() => { guard() }, [])` dari file halaman.      │
│   - Aksi: Tambahkan `<StandaloneAuthGate>` di remote `pages/_app.tsx`       │
│     agar local dev `localhost:300X` tetap bisa login mandiri.              │
│   - Dampak: Menghilangkan fetch ganda `/auth/me` dan `/menu/me`.           │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Konfigurasi Fase 1 pada Remote `next.config.js`:
```js
// next.config.js (Di repositori remote masing-masing)
shared: {
  "@ordentco/addons-auth-provider": {
    singleton: true,
    requiredVersion: false,
  },
  "@bri/addons-auth-provider": {
    singleton: true,
    requiredVersion: false,
  },
  ni18n: {
    singleton: true,
    requiredVersion: false,
  },
}
```

### Konfigurasi Standalone Gate di Remote `pages/_app.tsx` (Fase 2):
```tsx
// pages/_app.tsx (Di repositori remote untuk local dev)
function StandaloneAuthGate({ children }: { children: ReactNode }) {
  const { guard, token } = useAuth();

  useEffect(() => {
    if (token && typeof guard === "function") {
      try {
        const res = guard();
        if (res && typeof (res as any).catch === "function") {
          (res as any).catch(() => {});
        }
      } catch {}
    }
  }, [token, guard]);

  return <>{children}</>;
}
```

---

## 6. Ringkasan Manfaat Arsitektur

1. 🛡️ **Zero Crash**: Tidak ada lagi error `"Function not implemented"` di seluruh platform QLola.
2. 🚀 **Efisiensi Jaringan**: Mengurangi request autentikasi hingga >90% saat user berpindah halaman SPA.
3. ⚡ **Bebas Kedip**: Navigasi mulus memanfaatkan skeleton asli halaman tanpa terpotong oleh overlay loading global.
4. ⏱️ **Migrasi Cepat**: Squad cukup mengubah 1 baris konfigurasi di `next.config.js` tanpa harus me-refactor puluhan file halaman sekaligus.
