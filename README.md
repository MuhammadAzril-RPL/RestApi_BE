# Alur Program REST API

## Deskripsi

Program ini adalah sebuah REST API yang dibangun menggunakan bahasa pemrograman Go dan framework Gin. API ini mengambil data provinsi dari sumber eksternal (API publik), menyimpannya ke database MySQL, lalu menyajikannya dalam format JSON melalui endpoint.

## 1. Koneksi ke Database

query sql :
```sql
CREATE DATABASE IF NOT EXISTS wilayahs;
USE wilayahs;

CREATE TABLE  provinces (
    id INT AUTO_INCREMENT PRIMARY KEY,
    code VARCHAR(10) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL
);

```
Program pertaa-tama mencoba terhubung ke database MySQL menggunakan perintah:
```go
database, error := sql.Open("mysql", "Azril:Myboo5456@tcp(localhost:3307)/wilayahs")
```
Jika terjadi error saat koneksi, program akan dihentikan dengan:
```go
log.Fatal(error)
```
Koneksi ke database ditutup otomatis setelah program selesai dengan:
```go
defer database.Close()
```

## 2. Mengambil Data Provinsi dari API Publik

Program mengambil daftar provinsi dari API eksternal:
```go
response, error := http.Get("https://emsifa.github.io/api-wilayah-indonesia/api/provinces.json")
```
Jika gagal mengambil data, program akan dihentikan.
```go
	if error != nil {
		log.Fatal(error)
	}
	defer response.Body.Close()

```
Data ini kemudian disimpan dalam slice `provincesAPIList`:
```go
var provincesAPIList []ProvinceAPI
if error := json.NewDecoder(response.Body).Decode(&provincesAPIList); error != nil {
    log.Fatal(error)
}
```
## 3. Menyimpan Data ke Database

Setelah mendapatan data dari API, program akan menyimpan data tersebut ke dalam tabel `provinces` di MySQL:
```go
for _, province := range provincesAPIList {
    _, error := database.Exec("INSERT INTO provinces (code, name) VALUES (?, ?) ON DUPLICATE KEY UPDATE name = VALUES(name)", province.ID, province.Name)
}
```
Jika data provinsi sudah ada di database, maka hanya nama yang akan diperbarui.

## 4. Membuat Router dengan Gin

Program menggunakan framework Gin untuk menangani request HTTP. Router dibuat dengan:
```go
router := gin.Default()
```

## 5. Membuat Endpoint GET /province

Saat ada request ke `/province`, program akan mengambil data dari database:
```go
rows, error := database.Query("SELECT id, code, name FROM provinces")
```
Jika terjadi error saat query, program akan mengembalikan respons error 500.

Data yang berhasil diambil disimpan dalam slice `provincesList`:
```go
var provincesList []Province
for rows.Next() {
    var province Province
    if error := rows.Scan(&province.ID, &province.Code, &province.Name); error != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
        return
    }
    provincesList = append(provincesList, province)
}
```

## 6. Mengirim Respons dalam Format JSON

Setelah mendapatkan data, program akan mengembalikan response dalam format JSON:
```go
apiResponse := APIResponse{
    Status:  "success",
    Code:    200,
    Message: "Successfully get data",
    Data:    provincesList,
}
context.JSON(http.StatusOK, apiResponse)
```

```
## 7. Menjalankan Server

Server berjalan di `http://localhost:8000` dengan:
```go
if error := router.Run(":8000"); error != nil {
    log.Fatal(error)
}
```
Jika server berhasil dijalankan, program akan mencetak log:
```
Server berjalan di http://localhost:8000
```

## Struktur Data

- **ProvinceAPI**: Struktur untuk menampung data dari API eksternal.
- **Province**: Struktur untuk menampung data provinsi dari database.
- **APIResponse**: Struktur untuk membentuk respons API dalam format JSON.

