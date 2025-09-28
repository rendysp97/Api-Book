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
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL tidak ditemukan")
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

	fmt.Println("Success Connect DB")
}
