package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func OpenDB() *sql.DB {
	db, err := sql.Open("sqlite", "api.db?_pragma=foreign_keys(1)")
	if err != nil {
		log.Fatal("❌ Failed to connect to DB:", err)
	}
	_, err = db.Exec("PRAGMA journal_mode = WAL;")
	if err != nil {
		log.Fatal("❌ Failed to enable WAL mode:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("❌ Database unreachable:", err)
	}
	DB = db
	return db
}
