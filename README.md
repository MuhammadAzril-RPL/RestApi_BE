# Alur Program REST API

# Deskripsi 

Program ini adalah sebuah REST API yang dibangun pakai bahasa pemrograman Go dan framework Gin. API ini digunakan untuk ambil data provinsi dari database MySQL dan mengembalikannya dalam format JSON.

### Koneksi ke Database
Program pertama-tama coba terhubung ke database MySQL dengan `sql.Open()`.
dsn yang dipakai adalah:
- **Username**: Azril
- **Password**: Myboo5456
- **Host**: localhost
- **Port**: 3307
- **Database**: wilayahs
Kalau ada error saat koneksi, program langsung dihentikan dengan `log.Fatal(err)`.
Setelah selesai digunakan, koneksi ke database akan ditutup dengan `defer db.Close()`.

### Membuat Router dengan Gin
Program pakai framework Gin untuk tangani request HTTP. Router dibuat dengan `gin.Default()`.

### Membuat Endpoint GET /province
Ketika ada request ke endpoint `/province`, program akan ambil data dari tabel `provinces` di database. Data yang diambil meliputi `id`, `code`, dan `name`.

### Mengambil Data dari Database
Program jalankan query `SELECT id, code, name FROM provinces` untuk ambil data provinsi. Kalau ada error dalam query, program akan kembalikan response error HTTP 500. Data hasil query disimpan ke dalam slice `province` dengan tipe `Province`.

### Membuat Respons JSON
Setelah data berhasil diambil, program bungkus dalam struct `APIResponse`, yang berisi:
- **Status**: Status response (contohnya, "sukses").
- **Code**: Kode status HTTP (200).
- **Message**: Pesan yang menjelaskan hasil permintaan.
- **Data**: Slice yang berisi daftar provinsi dari database.
Response ini kemudian dikirim dalam format JSON pakai `c.JSON()`.

### Menjalankan Server
Server berjalan di port 8000 dengan `router.Run(":8000")`. Kalau server berhasil jalan, akan ada log `Server berjalan di http://localhost:8000`. Kalau ada error saat jalankan server, program akan berhenti dan tampilkan error.

## Struktur Data
- **Province**: Struct untuk representasikan satu provinsi dengan atribut `ID`, `Code`, dan `Name`.
- **APIResponse**: Struct yang digunakan untuk bungkus response API dalam format JSON dengan atribut `Status`, `Code`, `Message`, dan `Data`.
