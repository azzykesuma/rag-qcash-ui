# test case for qa creation

- **ID**: `antigravity-264f735e86bd9ffcbd4a01872521aa2f7b32013311c0d624c1b5fdfc5d1ec163`
- **Source Tool**: `antigravity`
- **Date**: `2026-08-31 13:16:38`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

test case for qa creation
https://docs.google.com/spreadsheets/d/1s4456K8zRx_NwFj5dBiWlUHXzosH1HrHxitC45yBSUw/edit?gid=1502635692#gid=1502635692
i have these repos, check the repos that only fall under my name (Azzy)
now i want to explain the scope the changes
1. i am removing ALL guard call on these repos
2. migrating  from ordent to bri

now i want to make a test case for my qa
can you 
1. generate a .md for each test cases for each repos that falls under my name according to the scope

since this is for qa, maybe we can make a less technical test case

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read_url_content**: 

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **find_by_name**: 

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **write_to_file**: 

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **write_to_file**: 

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 20: Assistant

Telah dibuat seluruh dokumen rencana pengujian (Test Case) berformat Markdown (`.md`) untuk **58 repositori** yang berada di bawah nama **Azzy** sesuai data pada Google Spreadsheet.

Semua file telah disimpan di dalam direktori proyek:
📁 [**`D:\work\qcash-ui\docs\qa-test-cases\`**](file:///D:/work/qcash-ui/docs/qa-test-cases/)

---

### 1. Rangkuman Ruang Lingkup Perubahan (Context for QA)

Untuk memudahkan tim QA memahami esensi perubahan tanpa terbebani istilah teknis pemrograman:
1. **Penghapusan Panggilan Otentikasi di Tiap Modul (*Removal of all `guard()` calls*)**:
   - **Sebelumnya**: Tiap modul (*remote MFE*) memanggil verifikasi sesi login (`guard()`) sendiri-sendiri saat dibuka, yang sering menyebabkan *flickering*, loading berulang, layar putih (*blank page*), atau sesi terputus saat berpindah-pindah menu.
   - **Sekarang**: Pengecekan sesi dipusatkan sepenuhnya di pintu gerbang portal utama ([`qcash-ui`](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui.md) via `HostAuthGate`). Tiap modul fitur tinggal menggunakan sesi login yang sudah diverifikasi oleh portal utama tanpa memvalidasi ulang.
2. **Migrasi Penyedia Otentikasi (Ordent ke BRI Provider)**:
   - Seluruh modul beralih dari pustaka otentikasi lama (`@ordentco/addons-auth-provider`) ke pustaka resmi BRI (`@bri/addons-auth-provider`).

---

### 2. Struktur Standar Test Case Tiap Modul (QA-Friendly)

Setiap file `.md` dirancang dengan pendekatan *user journey* dan *business flow*:
- **TC-01: Akses Awal Modul & Pemuatan Halaman (*Zero White Screen*)**: Memastikan halaman modul langsung terbuka mulus tanpa layar putih, kedipan berulang (*flicker*), atau *spinner* tak berujung.
- **TC-02: Verifikasi Hak Akses & Tombol Sesuai Wewenang (*RBAC*)**: Memastikan tombol tindakan (Tambah, Edit, Hapus, Kirim/Approve, Download) hanya muncul atau aktif sesuai wewenang role (*Maker*, *Approver*, *Viewer*, *Admin*).
- **TC-03: Pengoperasian Alur Bisnis Utama (*Core Workflow*)**: Menguji fungsi pencarian, filter tanggal, pengisian form, dan pembukaan modal konfirmasi.
- **TC-04: Navigasi Bolak-balik Antar Modul (*Cross-MFE Switching*)**: Berpindah dari Dashboard $\rightarrow$ Modul $\rightarrow$ Modul Lain $\rightarrow$ Kembali ke Modul tanpa terjadi pemutusan sesi (*logout* mendadak).
- **TC-05: Uji Ketahanan Muat Ulang Browser (*Page Refresh / F5*)**: Merefresh browser di dalam halaman modul; memastikan sesi tidak hilang dan halaman memuat ulang dengan normal.
- **TC-06: Penanganan Masa Sesi Habis (*Session Expiry & Re-login*)**: Saat waktu sesi habis, dialog resmi *Session Expired* muncul dengan rapi dan mengarahkan login ulang secara aman.
- **TC-07: Pemeriksaan Kerapian Jaringan DevTools (Opsional / Sanity Check)**: Memastikan tidak ada banjir panggilan berulang ke `/auth/me` atau `/menu/me` di tab Network.

---

### 3. Katalog Dokumen Test Case yang Dihasilkan

Dokumen indeks master lengkap tersedia di:  
📄 [**`D:\work\qcash-ui\docs\qa-test-cases\README.md`**](file:///D:/work/qcash-ui/docs/qa-test-cases/README.md)

#### A. Repositori Fondasi & Host (Platform Core)
| Nama Repositori | Peran & Ruang Lingkup | Tautan File Test Case |
|:---|:---|:---:|
| `qcash-ui` | **Host Shell Container** (Pusat otentikasi utama `HostAuthGate`) | [qcash-ui.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui.md) |
| `qcash-ui-header-footer` | **Navigasi Shell** (Menu navigasi, profil pengguna, logout) | [qcash-ui-header-footer.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-header-footer.md) |
| `qcash-ui-global-component` | **Shared UI Components** (Modal peringatan, form controls) | [qcash-ui-global-component.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-global-component.md) |
| `qcash-ui-dashboard` | **Beranda Utama** (Widget saldo, to-do list transaksi) | [qcash-ui-dashboard.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-dashboard.md) |
| `qcash-ui-base` | **Shared Layout Base MFE** (Layout wrapper & base UI) | [qcash-ui-base.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-base.md) |

#### B. Modul Transaksi & Pembayaran Utama
| Nama Repositori | Keterangan Fitur | Tautan File Test Case |
|:---|:---|:---:|
| `qcash-ui-fund-transfer` | Transfer Dana Domestik (BI-Fast, Online, RTGS, Kliring, Mass) | [qcash-ui-fund-transfer.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-fund-transfer.md) |
| `qcash-ui-bill-payment` | Pembayaran Tagihan Biller (PLN, Telkom, Pajak, PDAM, dll.) | [qcash-ui-bill-payment.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-bill-payment.md) |
| `qcash-ui-beneficiary-account-management` | Manajemen Rekening Tujuan Transfer | [qcash-ui-beneficiary-account-management.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-beneficiary-account-management.md) |
| `qcash-ui-account-statement` | Rekening Koran & Mutasi Transaksi | [qcash-ui-account-statement.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-account-statement.md) |
| `qcash-ui-account-summary` | Ringkasan Saldo Rekening Terkonsolidasi | [qcash-ui-account-summary.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-account-summary.md) |
| `qcash-ui-account-management` | Pengaturan & Akses Rekening Perusahaan | [qcash-ui-account-management.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-account-management.md) |
| `qcash-ui-create-briva` | Pembuatan & Administrasi Akun BRIVA | [qcash-ui-create-briva.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-create-briva.md) |
| `qcash-ui-ewallet-topup` | Top Up Dompet Digital (GoPay, OVO, DANA, ShopeePay) | [qcash-ui-ewallet-topup.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-ewallet-topup.md) |

#### C. Modul Monitoring, Supply Chain, Trade & Administrasi
| Nama Repositori | Keterangan Fitur & Catatan Khusus | Tautan File Test Case |
|:---|:---|:---:|
| `qcash-ui-bank-monitoring-transaction` | *Catatan Khusus QA: Uji intensif perpindahan antar fitur & tab filter* | [qcash-ui-bank-monitoring-transaction.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-bank-monitoring-transaction.md) |
| `qcash-ui-approval-signature` | Pengaturan Matriks Persetujuan & Tanda Tangan | [qcash-ui-approval-signature.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-approval-signature.md) |
| `qcash-ui-forex` | Transaksi Valas & Kontrak Kurs | [qcash-ui-forex.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-forex.md) |
| `qcash-ui-global-external-transfer` | Transfer Valas Internasional (SWIFT / Remittance) | [qcash-ui-global-external-transfer.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-global-external-transfer.md) |
| `qcash-ui-global-internal-transfer` | Transfer Internal Cabang Luar Negeri | [qcash-ui-global-internal-transfer.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-global-internal-transfer.md) |
| `qcash-ui-global-account-statement` | Rekening Koran Rekening Global Valas | [qcash-ui-global-account-statement.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-global-account-statement.md) |
| `qcash-ui-global-account-summary` | Ringkasan Saldo Valas Global | [qcash-ui-global-account-summary.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-global-account-summary.md) |
| `qcash-ui-df-pupuk-indonesia` | Distributor Financing Pupuk Indonesia | [qcash-ui-df-pupuk-indonesia.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-df-pupuk-indonesia.md) |
| `qcash-ui-do-pertamina` | Delivery Order Pertamina BBM/LPG | [qcash-ui-do-pertamina.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-do-pertamina.md) |
| `qcash-ui-do-pupuk-indonesia` | Delivery Order Pupuk Indonesia | [qcash-ui-do-pupuk-indonesia.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-do-pupuk-indonesia.md) |
| `qcash-ui-do-semen-indonesia` | Delivery Order Semen Indonesia (SIG) | [qcash-ui-do-semen-indonesia.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-do-semen-indonesia.md) |
| `qcash-ui-digi-guarantee` | Bank Garansi Digital | [qcash-ui-digi-guarantee.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-digi-guarantee.md) |
| `qcash-ui-advise-lc` | Advis Dokumen Letter of Credit | [qcash-ui-advise-lc.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-advise-lc.md) |
| `qcash-ui-amend-cancel-lc` | Perubahan & Pembatalan LC | [qcash-ui-amend-cancel-lc.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-amend-cancel-lc.md) |
| `qcash-ui-employee-data` | Data Karyawan & Master Payroll | [qcash-ui-employee-data.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-employee-data.md) |
| `qcash-ui-company-management` | Manajemen Profil & Cabang Perusahaan | [qcash-ui-company-management.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-company-management.md) |
| `qcash-ui-group-management` | Manajemen Grup Korporasi | [qcash-ui-group-management.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-group-management.md) |
| `qcash-ui-company-code-mapping` | Pemetaan Prefix & Kode Instansi | [qcash-ui-company-code-mapping.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-company-code-mapping.md) |
| `qcash-ui-auto-collection-registration` | Pendaftaran Mandat Auto Debet | [qcash-ui-auto-collection-registration.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-auto-collection-registration.md) |
| `qcash-ui-dashboard-auto-collection` | Dashboard & Analisis Auto Collection | [qcash-ui-dashboard-auto-collection.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-dashboard-auto-collection.md) |
| `qcash-ui-dashboard-dhe` | Monitoring Devisa Hasil Ekspor | [qcash-ui-dashboard-dhe.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-dashboard-dhe.md) |
| `qcash-ui-dashboard-merchant` | Dashboard & Rekapitulasi Merchant EDC/QRIS | [qcash-ui-dashboard-merchant.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-dashboard-merchant.md) |
| `qcash-ui-credit-card` | Informasi Kartu Kredit Korporat | [qcash-ui-credit-card.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-credit-card.md) |
| `qcash-ui-cashcard` | Manajemen Kartu Cashcard | [qcash-ui-cashcard.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-cashcard.md) |
| `qcash-ui-deposito` | Pembukaan & Pengelolaan Deposito | [qcash-ui-deposito.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-deposito.md) |
| `qcash-ui-bond-price` | Harga Kuotasi Obligasi Pasar | [qcash-ui-bond-price.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-bond-price.md) |
| `qcash-ui-announcement-management` | Pengumuman & Informasi Portal | [qcash-ui-announcement-management.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-announcement-management.md) |
| `qcash-ui-complaint` | Pengaduan Nasabah Korporat | [qcash-ui-complaint.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-complaint.md) |
| `qcash-ui-content-management-system` | Manajemen Banner & Konten CMS | [qcash-ui-content-management-system.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-content-management-system.md) |
| `qcash-ui-cut-off-management` | Manajemen Waktu Cut-Off Transaksi | [qcash-ui-cut-off-management.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-cut-off-management.md) |
| `qcash-ui-helpdesk-portal` | Portal Tiket Helpdesk | [qcash-ui-helpdesk-portal.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-helpdesk-portal.md) |
| `qcash-ui-holiday-management` | Kalender Operasional Hari Libur | [qcash-ui-holiday-management.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-holiday-management.md) |
| `qcash-ui-account-receivable` | Piutang Usaha & Pelunasan Invoice | [qcash-ui-account-receivable.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-account-receivable.md) |
| `qcash-ui-account-registration` | Pendaftaran Rekening Baru | [qcash-ui-account-registration.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-account-registration.md) |
| `qcash-ui-account-validation` | Validasi Nama & Rekening Penerima | [qcash-ui-account-validation.md](file:///D:/work/qcash-ui/docs/qa-test-cases/qcash-ui-account-validation.md) |

#### D. Modul Non-Aktif / Statis / Legacy (Status N/A)
Dokumen test case untuk repositori ini telah ditandai secara jelas berstatus **Statis / Non-Aktif** sehingga tim QA tidak perlu membuang waktu melakukan pengetesan terpisah:
- `qcash-ui-bgn.md` (Aplikasi BGN mandiri statis)
- `qcash-ui-bifast.md` (Legacy, sudah digabung ke Fund Transfer)
- `qcash-ui-bricams-coming-soon.md` (Halaman landing statis)
- `qcash-ui-bricams-redirect.md` (Pengalihan URL)
- `qcash-ui-creation.md` (Template stub internal)
- `qcash-ui-deposito-online.md` (Repo kosong)
- `qcash-ui-ecosystem.md` (Etalase portal statis)
- `qcash-ui-external-fund-transfer.md` (Legacy stub)
- `qcash-ui-history-external-fund-transfer.md` (Legacy stub)
- `qcash-ui-ift.md` (Legacy stub)

Dokumen-dokumen ini siap dibagikan langsung kepada tim QA sebagai panduan pengujian rilis prestage maupun pilot.

---

## Turn 21: User

can you merge them into one, and make a .xls file?

---

## Turn 22: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 23: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 24: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **manage_task**: 

</details>



---

## Turn 25: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **manage_task**: 

</details>



---

## Turn 26: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **write_to_file**: 

</details>



---

## Turn 27: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 28: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 29: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 30: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 31: Assistant

Kedua kebutuhan Anda telah selesai dibuat dan diverifikasi:

1. **Merged Markdown (`.md`)**: Seluruh test case dari ke-58 repositori telah digabungkan menjadi satu dokumen terpadu lengkap dengan Daftar Isi navigasi langsung (anchor links).
2. **File Spreadsheet Excel (`.xls`)**: Dokumen spreadsheet dengan format XML Spreadsheet 2003 yang didukung secara native oleh Microsoft Excel, lengkap dengan styling warna korporat BRI, pemisahan kolom langkah uji, hasil yang diharapkan, checklist pass/fail, dan auto-wrap text.

---

### Lokasi File yang Dihasilkan

| Format | Nama File & Lokasi | Ukuran | Keterangan |
|:---|:---|:---:|:---|
| **Excel (`.xls`)** | [**`D:\work\qcash-ui\docs\qa-test-cases\QA_TEST_CASES_ALL_REPOS.xls`**](file:///D:/work/qcash-ui/docs/qa-test-cases/QA_TEST_CASES_ALL_REPOS.xls) | ~509 KB | Siap dibuka langsung di Microsoft Excel dengan 2 lembar kerja (*sheets*). |
| **Merged Markdown (`.md`)** | [**`D:\work\qcash-ui\docs\qa-test-cases\QA_TEST_CASES_ALL_REPOS.md`**](file:///D:/work/qcash-ui/docs/qa-test-cases/QA_TEST_CASES_ALL_REPOS.md) | ~435 KB | 1 file Markdown utuh mencakup seluruh 58 modul, 7 skenario per modul aktif, dan daftar isi. |
| **Indeks Katalog** | [**`D:\work\qcash-ui\docs\qa-test-cases\README.md`**](file:///D:/work/qcash-ui/docs/qa-test-cases/README.md) | ~16 KB | Katalog ringkas penjelas ruang lingkup bagi QA. |

---

### Struktur File Excel (`.xls`)

File Excel ini dirancang khusus agar tim QA dapat langsung melakukan eksekusi uji dan menandai status `PASS` / `FAIL` dengan rapi:

#### 1. Lembar Kerja 1: `Ringkasan Modul` (Summary)
Daftar inventaris seluruh 58 repositori yang mencakup:
- **No & Nama Repositori** (format kode monospaced)
- **Nama Fitur / Modul Bisnis** (misal: *Transfer Dana Domestik*, *Monitoring Transaksi Bank*)
- **Kategori Fitur** (*Platform Core*, *Transfer & Payments*, *Trade Finance*, dll.)
- **Status Modul** (Label hijau untuk `Aktif (Perlu Uji)`, label abu-abu untuk `Statis / Non-Aktif`)
- **Target Wewenang Role** (*Maker*, *Approver*, *Releaser*, *Admin*)
- **PIC Developer** (`Azzy`)
- **Tautan PR Dev** (Link Bitbucket dev untuk PR yang tersedia)
- **Catatan Khusus QA** (Contoh: Catatan pengetesan perpindahan tab filter pada *Bank Monitoring Transaction*)
- **Kolom Hasil Akhir QA** (`UNTESTED` / `PASS` / `FAIL`)

#### 2. Lembar Kerja 2: `Detail Skenario Uji` (Detailed Test Cases - 352 Baris Skenario)
Setiap repositori aktif dipecah menjadi **7 skenario pengujian fungsional non-teknis**:
1. **TC-01: Akses Awal Modul & Pemuatan Halaman**: Memastikan tidak ada *white screen*, tidak ada kedipan (*flicker*), dan tidak ada *infinite spinner*.
2. **TC-02: Verifikasi Hak Akses & Tombol Sesuai Role (RBAC)**: Memastikan tombol aksi (Tambah, Filter, Download, Edit, Approve) hanya tampil sesuai wewenang akun yang login.
3. **TC-03: Pengoperasian Fitur Utama & Interaksi Data**: Menjalankan alur bisnis spesifik modul terkait (input form, filter tabel, buka detail modal).
4. **TC-04: Navigasi Bolak-balik Antar Modul**: Berpindah modul ke Dashboard $\rightarrow$ Modul Lain $\rightarrow$ Kembali ke Modul tanpa terjadi pemutusan sesi.
5. **TC-05: Uji Ketahanan Muat Ulang Browser (Refresh / F5)**: Refresh halaman di dalam modul; memastikan sesi tidak ter-reset atau terlempar ke layar login.
6. **TC-06: Penanganan Sesi Habis (Session Expiry & Re-login)**: Memastikan pop-up resmi *Session Expired* muncul dengan benar saat sesi kadaluarsa.
7. **TC-07: Pemeriksaan Kerapian Jaringan DevTools (Sanity Check)**: Memastikan tidak ada lonjakan panggilan berulang ke `/auth/me` atau `/menu/me` di tab Network.

---

### Cara Penggunaan oleh Tim QA

1. **Menggunakan Excel**:
   - Buka file [**`QA_TEST_CASES_ALL_REPOS.xls`**](file:///D:/work/qcash-ui/docs/qa-test-cases/QA_TEST_CASES_ALL_REPOS.xls) menggunakan Microsoft Excel / LibreOffice / Google Sheets.
   - Tim QA dapat memfilter baris per nama modul atau kategori pada tab **`Detail Skenario Uji`**.
   - Isi kolom **Status Uji** dengan `PASS` / `FAIL` dan tambahkan catatan bug/temuan di kolom **Catatan / Temuan QA**.
2. **Menggunakan Markdown Terpadu**:
   - Buka file [**`QA_TEST_CASES_ALL_REPOS.md`**](file:///D:/work/qcash-ui/docs/qa-test-cases/QA_TEST_CASES_ALL_REPOS.md) di VS Code, GitHub/Bitbucket preview, atau markdown viewer pilihan.
   - Gunakan link pada bagian Daftar Isi di awal dokumen untuk melompat langsung ke repositori yang sedang diuji.

---

