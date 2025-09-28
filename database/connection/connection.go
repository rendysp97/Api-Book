package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDB() {
	// Ambil DATABASE_URL dari environment (Railway)
	databaseUrl := os.Getenv("DATABASE_URL")

	// Kalau DATABASE_URL kosong, fallback ke database lokal
	if databaseUrl == "" {
		fmt.Println("DATABASE_URL tidak ditemukan, pakai lokal")
		databaseUrl = "postgresql://postgres:postgres@localhost:5432/book_db" // ganti sesuai lokalmu
	}

	var err error
	Db, err = sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal("Gagal buka koneksi:", err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal("Database tidak merespon:", err)
	}

	fmt.Println("Success Connect DB:", databaseUrl)
}
