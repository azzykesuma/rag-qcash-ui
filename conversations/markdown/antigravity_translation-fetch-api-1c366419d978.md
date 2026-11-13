# check this screeenshot. it seems there are too many translation fet...

- **ID**: `antigravity-e46fd2f8abe4a178c0ee597bab78afd7a373374d6fdcff75e27a1c366419d978`
- **Source Tool**: `antigravity`
- **Date**: `2026-11-11 09:13:33`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

check this screeenshot. it seems there are too many translation fetch api, can you investigate why it's like this?

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 21: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 22: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 23: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 24: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 25: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 26: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 27: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 28: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 29: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 30: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 31: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 32: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 33: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 34: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 35: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_file**: 

</details>



---

## Turn 36: Assistant

Berdasarkan screenshot DevTools Network Anda (`internal-service.example.com/v2/landing-page`), ada **14 request translation** yang terjadi secara bersamaan.

Penyebab utamanya adalah **3 layer pemanggilan translation yang berjalan paralel tanpa koordinasi dan deduplikasi** antara Host (`qcash-ui`) dan Remote MFE (`qcash-ui-landing-page`).

---

### Pemetaan 14 Request pada Screenshot

| No | Request | Status & Tipe | Initiator | Asal Muasal / Penyebab |
|---|---|---|---|---|
| **1-3** | `translation.json`, `header-footer`, `common` (`lang=en`) | 304 / 200 (fetch) | `26954...` (Host Bundle) | **Host ni18n initial hydration**: Saat aplikasi pertama kali hydrate, `ni18n` membaca default/fallback language (`en`). |
| **4-6** | `translation.json`, `header-footer`, `common` (`lang=id`) | 304 / 200 (fetch) | `26954...` (Host Bundle) | **Host language switch ke "id"**: Begitu [`ModuleFederationProvider`](file:///C:/Work/qcash-ui/components/federation/provider.tsx#L81-L85) me-mount `useSyncLanguage(localStorage.getItem("locale"))`, bahasa berubah dari `en` ke `id`, memicu `ni18n` me-load ulang seluruh namespace `id`. |
| **7-9** | `landing-page`, `common`, `header-footer` (`lang=id`) | 200 (fetch) | `5922...` (Host MFE Loader) | **Host `useModuleLoading`**: Saat route `/landing-page` aktif, [`dynamicLoadLocales`](file:///C:/Work/qcash-ui/services/federation/utils/i18n.ts#L164) dipanggil. Fungsi ini **meng-hardcode** penambahan `"common"` dan `"header-footer"` ke dalam namespace modul MFE, sehingga men-download ulang `common` dan `header-footer` yang tadi sudah di-fetch oleh `ni18n`. |
| **10** | `landing-page` / `common` (`lang=en`) | 200 (fetch) | `26954...` (Host Bundle) | **`reloadResources` tanpa spesifikasi bahasa**: Pada [`useModuleLoading.ts#L133`](file:///C:/Work/qcash-ui/components/federation/hooks/useModuleLoading.ts#L133), ada pemanggilan `i18n.reloadResources(undefined, ["common", "header-footer"])`. Parameter `undefined` menyebabkan i18next me-load ulang semua bahasa (termasuk fallback `en`). |
| **11-14** | `common`, `landing-page` (`lang=id`) | 200 (**xhr**) | `_federation_expose_...` | **Remote MFE Axios calls**: Remote MFE memiliki hook sendiri ([`useLanguageSync.ts`](file:///C:/Work/qcash-ui-landing-page/hooks/useLanguageSync.ts#L143-L190)) yang memanggil API via Axios (XHR). Karena hook ini dipanggil di **3 komponen berbeda sekaligus**, request terduplikasi ganda. |

---

### Akar Masalah & Rekomendasi Solusi

#### 1. Remote MFE memanggil `useLanguageSync()` sebanyak 3 kali secara paralel
* **Penyebab**: Di repository `qcash-ui-landing-page`, hook `useLanguageSync()` dipanggil di:
  1. [`hooks/useLandingPageContent.ts`](file:///C:/Work/qcash-ui-landing-page/hooks/useLandingPageContent.ts#L7) (dipakai di `LandingPageContent`)
  2. [`containers/landing-page.tsx`](file:///C:/Work/qcash-ui-landing-page/containers/landing-page.tsx#L26) (`LandingPageQlola`)
  3. [`components/newskin/navbar/navbar-newskin.tsx`](file:///C:/Work/qcash-ui-landing-page/components/newskin/navbar/navbar-newskin.tsx#L47) (`NavbarNewSkin`)
  Karena ketiganya me-mount pada waktu yang sama dan memiliki state/instance terpisah, masing-masing menembak Axios `getTranslations` & `getCommonTranslations` secara bersamaan (menghasilkan request 11–14).
* **Solusi**: 
  - Hapus pemanggilan `useLanguageSync()` dari `LandingPageQlola` dan `NavbarNewSkin`. Cukup panggil **sekali saja** di level parent/hook utama (`useLandingPageContent`).
  - Tambahkan in-flight promise deduplication atau rely pada singleton i18n host jika running di dalam host MFE.

#### 2. Redundansi Namespace di Host MFE Loader (`dynamicLoadLocales`)
* **Penyebab**: Di [`services/federation/utils/i18n.ts#L164`](file:///C:/Work/qcash-ui/services/federation/utils/i18n.ts#L164):
  ```ts
  const nss = Array.from(new Set([...requestedNs, "common", "header-footer"]));
  ```
  Host selalu memaksa me-fetch `"common"` dan `"header-footer"` setiap kali memuat locales modul MFE, padahal kedua namespace ini sudah dimuat oleh `ni18n`.
* **Solusi**: Hanya fetch `feature.i18n` (misal `["landing-page"]`) dan hindari menambah `"common"` & `"header-footer"` jika namespace tersebut sudah di-*handle* oleh host `ni18n`.

#### 3. Pemanggilan `i18n.reloadResources(undefined, ...)`
* **Penyebab**: Di [`components/federation/hooks/useModuleLoading.ts#L133`](file:///C:/Work/qcash-ui/components/federation/hooks/useModuleLoading.ts#L133):
  ```ts
  i18n?.reloadResources(undefined, ["common", "header-footer"]);
  ```
  Ini dieksekusi setiap kali rute dimuat, memicu reload untuk semua bahasa fallback (`en` dan `id`).
* **Solusi**: Batasi pemanggilan `reloadResources` hanya saat benar-benar ada perubahan bahasa atau sertakan bahasa aktif saja (`[i18n.language]`).

#### 4. Flip Bahasa Awal (`en` ➔ `id`) Saat Hydration
* **Penyebab**: SSR/initial render me-load resource bahasa default `en`, lalu segera setelah mount di browser membaca `localStorage.getItem("locale") = "id"` dan me-load ulang semua resource untuk `id`.
* **Solusi**: Pastikan deteksi bahasa memanfaatkan cookie/header pada saat SSR atau `ni18n` detection config agar langsung menginisialisasi bahasa yang sesuai sejak awal render.

---

