# Golang CRUD API dengan Echo dan SQLite

Ini adalah API sederhana yang mengimplementasikan operasi CRUD (Create, Read, Update, Delete) menggunakan framework [Echo](https://echo.labstack.com/) di Golang dengan database SQLite. API ini mendukung operasi CRUD untuk resource `Product`, termasuk membuat, membaca, memperbarui, dan menghapus produk.

## Persyaratan

- [Golang](https://golang.org/doc/install) (versi 1.16 ke atas)
- [Framework Echo](https://echo.labstack.com/)
- [SQLite3](https://www.sqlite.org/download.html) atau [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite)
- Postman (atau alat testing API lainnya)

## Instalasi dan Setup

1. **Clone repository**:

   ```bash
   git clone https://github.com/username/repository.git
   cd repository
   ```

2. **Install dependencies**:

   Pastikan Go modules diaktifkan, lalu jalankan:

   ```bash
   go mod tidy
   ```

3. **Jalankan aplikasi**:

   Setelah dependensi terpasang, Anda bisa menjalankan server Go dengan perintah:

   ```bash
   go run main.go
   ```

   Server akan berjalan di `http://localhost:8080`.

## Endpoint API

### 1. Membuat Produk

- **URL**: `/products`
- **Metode**: `POST`
- **Body**: JSON

  ```json
  {
    "name": "Nama Produk",
    "price": 1000
  }
  ```

- **Response**: `201 Created`

  ```json
  {
    "id": 1,
    "name": "Nama Produk",
    "price": 1000
  }
  ```

### 2. Mendapatkan Semua Produk

- **URL**: `/products`
- **Metode**: `GET`
- **Response**: `200 OK`

  ```json
  [
    {
      "id": 1,
      "name": "Nama Produk",
      "price": 1000
    },
    {
      "id": 2,
      "name": "Produk Lain",
      "price": 500
    }
  ]
  ```

### 3. Mendapatkan Produk Berdasarkan ID

- **URL**: `/products/{id}`
- **Metode**: `GET`
- **Response**: `200 OK`

  ```json
  {
    "id": 1,
    "name": "Nama Produk",
    "price": 1000
  }
  ```

### 4. Mengubah Produk

- **URL**: `/products/{id}`
- **Metode**: `PUT`
- **Body**: JSON

  ```json
  {
    "name": "Nama Produk Baru",
    "price": 1500
  }
  ```

- **Response**: `200 OK`

  ```json
  {
    "id": 1,
    "name": "Nama Produk Baru",
    "price": 1500
  }
  ```

### 5. Menghapus Produk

- **URL**: `/products/{id}`
- **Metode**: `DELETE`
- **Response**: `200 OK`

  ```json
  "Product deleted successfully"
  ```

## Database

Proyek ini menggunakan SQLite sebagai database. Secara default, file database (`database.db`) akan dibuat di root direktori proyek.

### Cara Melihat atau Memodifikasi Database:
- Anda bisa menggunakan alat seperti **[DB Browser for SQLite](https://sqlitebrowser.org/)** untuk melihat dan mengelola database SQLite.
- File database akan dibuat secara otomatis saat Anda menjalankan aplikasi jika belum ada.

### Skema

Database secara otomatis membuat tabel ketika aplikasi berjalan. Tabel yang dibuat adalah `products` dengan field sebagai berikut:

- `id`: INTEGER PRIMARY KEY AUTOINCREMENT
- `name`: TEXT
- `price`: INTEGER

## Cara Menguji API Menggunakan Postman

Anda dapat menguji API ini menggunakan [Postman](https://www.postman.com/) atau alat sejenis dengan langkah-langkah berikut:

1. **Jalankan server**: Jalankan `go run .` di direktori proyek Anda.
2. **Gunakan Postman**: Buka Postman dan gunakan pengaturan berikut:
   - **POST**: `http://localhost:8080/products` (untuk membuat produk)
   - **GET**: `http://localhost:8080/products` (untuk mendapatkan semua produk)
   - **GET**: `http://localhost:8080/products/{id}` (untuk mendapatkan produk tertentu berdasarkan ID)
   - **PUT**: `http://localhost:8080/products/{id}` (untuk memperbarui produk berdasarkan ID)
   - **DELETE**: `http://localhost:8080/products/{id}` (untuk menghapus produk berdasarkan ID)


### Penjelasan:

- **Persyaratan**: Bagian ini menjelaskan kebutuhan sistem dan alat yang harus diinstal untuk menjalankan proyek.
- **Instalasi dan Setup**: Langkah-langkah untuk menyiapkan dan menjalankan aplikasi di mesin lokal.
- **Endpoint API**: Penjelasan tentang endpoint CRUD, metode HTTP yang digunakan, dan contoh JSON request/response.
- **Database**: Memberikan informasi tentang penggunaan SQLite sebagai database, serta cara mengakses dan memodifikasinya.
- **Testing dengan Postman**: Instruksi tentang cara menguji API dengan Postman atau alat serupa.
- **Lisensi**: Menjelaskan bahwa proyek dilisensikan di bawah lisensi MIT.
