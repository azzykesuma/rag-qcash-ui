# Panduan Lengkap Host Auth Gate, Matriks Skenario & Rincian Perubahan Kode (Code Changes)

Dokumen ini ditujukan untuk dibagikan kepada rekan tim / squad lead sebagai panduan implementasi **Host Auth Gate**, alur **Local Development**, serta contoh perubahan kode (**Before vs After**) pada Host Shell dan Remote MFE.

---

## 1. Alur Standar Local Development (Host + Remote)

Untuk melakukan development fitur di lokal, developer **selalu menjalankan 2 repository secara bersamaan**:

```
┌─────────────────────────────────────────────────────────────────────────────────┐
│ 1. Jalankan HOST SHELL (`qcash-ui`)               ──► http://localhost:3000     │
│ 2. Jalankan REMOTE FITUR (misal: `fund-transfer`) ──► http://localhost:3001     │
│ 3. Buka Browser di HOST:                          ──► http://localhost:3000/... │
└─────────────────────────────────────────────────────────────────────────────────┘
```

### Mengapa Alurnya Selalu Host + Remote?
1. **Host adalah Pemilik Autentikasi**: Host Shell (`localhost:3000`) mengelola token, login session, header/footer, dan navigasi global melalui `<HostAuthGate />`.
2. **Remote adalah Pure Feature Consumer**: Remote (`localhost:3001`) hanya menyajikan komponen UI dan logika bisnis fiturnya.
3. **Module Federation Menggabungkan Keduanya**: Host me-load komponen dari remote lokal secara realtime dengan hot-reload.

---

## 2. Apa itu Host Auth Gate (`<HostAuthGate />`)?

`<HostAuthGate />` adalah gerbang autentikasi terpusat yang dipasang di **Host Shell (`qcash-ui/pages/_app.tsx`)**.

### Cara Kerjanya:
1. **Membaca Token Otomatis**: Mengambil `access-token` dari `localStorage` dan memasukkannya ke React Context.
2. **Validasi 1 Kali Saja**: Menjalankan fungsi `guard()` **tepat 1 kali per sesi login** untuk mengambil data user (`/auth/me`) dan hak akses menu (`/menu/me`).
3. **Penanganan Otomatis Error 401**: Jika token kadaluarsa/invalid, sistem otomatis menghapus token dan me-redirect user ke halaman login `/landing-page`.
4. **Bebas Kedip (No Blink)**: Langsung menampilkan halaman dan skeleton bawaan tanpa overlay penutup yang mengganggu.

---

## 3. Matriks Semua Kemungkinan Skenario

| No | Singleton di `next.config.js` | Versi Provider (Host vs Remote) | Pemanggilan `guard()` di Halaman Remote | Status & Perilaku Aplikasi | Jumlah Request `/auth/me` & `/menu/me` |
| :---: | :---: | :---: | :---: | :--- | :---: |
| **1** | ❌ Tidak (atau Commented) | ❌ Beda Versi | ✅ Masih Ada | 💥 **CRASH: `Error: Function not implemented`**.<br>Remote membuat context terisolasi yang belum diinisialisasi. | Banyak (Gagal/Error) |
| **2** | ❌ Tidak | ✅ Versi Sama | ✅ Masih Ada | ⚠️ **Infinite Loading / Stuck**.<br>Remote mencoba memanggil guard lokal, tetapi context terpisah sehingga state tidak sinkron dengan Host. | 2x per halaman |
| **3** | ❌ Tidak | Bebas | ❌ Sudah Dihapus | 🟡 **Aman dari crash**, tapi `useAuth()` di remote menghasilkan nilai kosong/default karena tidak terhubung ke context Host. | 1x (Hanya dari Host) |
| **4** | ✅ **`singleton: true`** | ❌ Beda Versi | ✅ Masih Ada | 🟢 **Aman / Berjalan Normal**.<br>Meskipun versi beda, Webpack Module Federation mencoba menyatukan context instance. | 2x saat halaman dibuka (1x Host + 1x Remote) |
| **5** | ✅ **`singleton: true`** | ✅ **Versi Sama** | ✅ Masih Ada | 🟢 **Stabil & Tidak Crash**.<br>Remote membaca context Host secara sempurna, namun masih ada request ganda dari `useEffect` remote. | 2x saat halaman dibuka (1x Host + 1x Remote) |
| **6** | ✅ **`singleton: true`** | ✅ **Versi Sama** | ❌ **Sudah Dihapus** | 🚀 **SUPER OPTIMAL (Target Akhir)**.<br>Bebas crash, instan saat ganti menu, dan **hemat request jaringan >90%**. | **Cuma 1x untuk seluruh sesi login!** |

---

## 4. Rincian Perubahan Kode (Code Changes)

Berikut adalah panduan baris kode yang perlu diubah:

### A. Perubahan pada SISI HOST SHELL (`qcash-ui`)

#### 1. File `pages/_app.tsx` di Host
Tambahkan `<HostAuthGate>` membungkus komponen aplikasi:
```diff
  // qcash-ui/pages/_app.tsx
+ import { HostAuthGate } from "@/components/providers/HostAuthGate";

  export default function App({ Component, pageProps }: AppProps) {
    return (
      <BriProviders>
        <OrdentProviders>
+         <HostAuthGate>
            <AuthBridgeSync />
            <Component {...pageProps} />
+         </HostAuthGate>
        </OrdentProviders>
      </BriProviders>
    );
  }
```

---

### B. Perubahan pada SISI REPO REMOTE SQUAD (Misal: `qcash-ui-fund-transfer`, `qcash-ui-credit-card`, dll.)

#### 1. File `next.config.js` di Remote (Langkah 1 — Wajib)
Tambahkan `singleton: true` pada `shared`:
```diff
  // next.config.js (Di repositori remote)
  shared: {
+   "@ordentco/addons-auth-provider": {
+     singleton: true,
+     requiredVersion: false,
+   },
+   "@bri/addons-auth-provider": {
+     singleton: true,
+     requiredVersion: false,
+   },
    ni18n: {
      singleton: true,
      requiredVersion: false,
    },
  }
```
👉 *Efek: Langsung menghentikan semua crash `"Function not implemented"`.*

---

#### 2. File Halaman `pages/**/*.tsx` di Remote (Langkah 2 — Bersihkan Guard)

##### **Before (Pola Lama — Menimbulkan Fetch Ganda):**
```tsx
// ❌ POLA LAMA
import { useEffect } from "react";
import { useAuth } from "@ordentco/addons-auth-provider";

export default function FundTransferPage() {
  const { guard, isAuthoritiesReady, productAuthorities } = useAuth();

  // ⚠️ Memanggil guard di setiap halaman menyebabkan fetch ganda ke /auth/me & /menu/me
  useEffect(() => {
    guard();
  }, []);

  if (!isAuthoritiesReady) return <LoadingOverlay />;

  return <div>Konten Fitur</div>;
}
```

##### **After (Pola Baru — Bersih & Super Cepat):**
```tsx
// ✅ POLA BARU (Direkomendasikan)
import { useAuth } from "@ordentco/addons-auth-provider";

export default function FundTransferPage() {
  const { isAuthoritiesReady, productAuthorities } = useAuth();

  // Gunakan optional chaining (?.) untuk keamanan pembacaan hak akses
  const canAccess = productAuthorities?.INTERNAL_FUND_TRANSFER?.anyAuthority;

  if (!isAuthoritiesReady) return <LoadingOverlay />;
  if (!canAccess) return <UnauthorizedDialog />;

  return <div>Konten Fitur</div>;
}
```

---

## 5. Ringkasan Manfaat untuk Tim

1. **Bebas Crash**: Tidak ada lagi error `"Function not implemented"` saat remote di-load.
2. **Hemat Traffic Jaringan**: Menghilangkan request `/auth/me` dan `/menu/me` yang berulang-ulang di setiap navigasi halaman.
3. **Penerapan Sangat Cepat**: Squad hanya perlu mengubah 1 file (`next.config.js`) untuk fase awal, kemudian membersihkan `guard()` di file halaman.
