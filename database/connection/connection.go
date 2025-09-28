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
	// Ambil DATABASE_URL dari env (Railway)
	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl != "" {
		fmt.Println("Menggunakan DATABASE_URL dari Railway")
	} else {
		// Fallback ke lokal (pakai IPv4 untuk menghindari masalah IPv6)
		databaseUrl = "postgresql://postgres:postgres@127.0.0.1:5432/book_db"
		fmt.Println("Menggunakan database lokal:", databaseUrl)
	}

	var err error
	Db, err = sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal("Gagal buka koneksi:", err)
	}

	// Ping untuk memastikan database merespon
	err = Db.Ping()
	if err != nil {
		log.Fatal("Database tidak merespon:", err)
	}

	fmt.Println("Success Connect DB")
}
