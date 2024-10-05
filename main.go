package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	_ "modernc.org/sqlite" // Menggunakan driver pure Go
)

var db *sql.DB

func initDB() {
	var err error

	// Membuka koneksi ke file SQLite database
	db, err = sql.Open("sqlite", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	// Membuat tabel 'products' jika belum ada
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS products (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        price INTEGER
    );`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Inisialisasi database SQLite
	initDB()

	// Membuat instance echo
	e := echo.New()

	// Mendaftarkan route
	RegisterRoutes(e)

	// Menjalankan server pada port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
