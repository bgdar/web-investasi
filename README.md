# web inverstasi

### Tech stack

<p align="center">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=golang,postgres,html,js,tailwind&perline=4" />
  </a>
</p>

**Back**

- `GIN` : web framework dari golang
  **Front**
- `Tailwindcss` : untuk style CSS
- `fontawesome` : font yang di guanakan di web ini
- `pq` : conector untuk koneksi ke database PogressSql
- `jwt` : untuk Login user mengguanakn Json web Token
  client : menggunaakn Cookies
  > cek configurasi database di /config/config.go

### Controller

- base.go : untuk menampilkan atau tampilan utama aplikasi
  1. dashboard
  2. block atau about
- product.go : api untuk trasaksi product,product nya :
  . `robot` :
- transaksi : untuk halaman trasaksi terjadi atau seperti payment gateway

### seeder

rencanyan untuk data awal yang akan di isi atau structure dari APP

> untuk awal pengembangan atau deployment

- `table` : table table yang akan di generate di awal

### Views

- /Component/ : berisi component component yang akan di gunakan di setiap VIEWS

pada **Views** ada folder **Componet** di sini menyimpan component yang bsai di gunakan di views

- header.html : ini file yang berisi management di <head>....</head> tag

### Color app

| Kategori                | Nama Warna | Kode      | Kesan / Fungsi                    |
| ----------------------- | ---------- | --------- | --------------------------------- |
| **Utama & Netral**      | Slate Gray | `#6B7280` | Abu-abu elegan                    |
|                         | Cool Gray  | `#9CA3AF` | Netral lembut                     |
|                         | Charcoal   | `#374151` | Tegas tapi tidak mencolok         |
|                         | Gainsboro  | `#DADADA` | Garis batas halus                 |
|                         | Jet Black  | `#111827` | Kontras kuat untuk teks           |
| **Tambahan / Aksen**    | Steel Blue | `#4682B4` | Aksen tenang dan profesional      |
|                         | Teal Gray  | `#5F9EA0` | Biru kehijauan lembut             |
|                         | Soft Beige | `#E8E6E3` | Memberi kehangatan ringan         |
|                         | Warm Taupe | `#B2A59F` | Kesan premium dan seimbang        |
|                         | Silver     | `#C0C0C0` | Penegas visual netral             |
| **Status / Notifikasi** | Info       | `#3B82F6` | Biru lembut (informasi / pesan)   |
|                         | Success    | `#22C55E` | Hijau lembut (berhasil / positif) |
|                         | Warning    | `#FACC15` | Kuning cerah (peringatan)         |
|                         | Error      | `#EF4444` | Merah tegas tapi tidak mencolok   |
|                         | Neutral    | `#9CA3AF` | Untuk status abu-abu netral       |

example penggunana
| Elemen UI | Warna Disarankan |
| ------------------ | ------------------------ |
| Background utama | `#fcfcfc` |
| Teks utama | `#111827` (Jet Black) |
| Teks sekunder | `#6B7280` (Slate Gray) |
| Border / Divider | `#DADADA` (Gainsboro) |
| Card / Panel | `#E8E6E3` || `#F3F4F6`|
| Tombol utama | `#374151` (Charcoal) |
| Hover tombol | `#4682B4` (Steel Blue) |
| Icon / Accent | `#5F9EA0` (Teal Gray) |
| Link aktif | `#3B82F6` (Info Blue) |
| Notifikasi sukses | `#22C55E` |
| Notifikasi error | `#EF4444` |
| Notifikasi warning | `#FACC15` |
| Notifikasi info | `#3B82F6` |

### database

![Structure table database]("./structure table.png")
<br>
Akses database PogressSql

```bash
# linux  ( dengan user default postgres )
psql -U postgres -W

```
