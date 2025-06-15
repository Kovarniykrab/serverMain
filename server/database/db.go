package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

var db *sql.DB

const dbFile = "tableOne"

// InitDB инициализирует подключение
func InitDB() (*sql.DB, error) {

	var err error

	db, err = sql.Open("postgress", dbFile)
	if err != nil {
		return nil, fmt.Errorf("error oper database: %v", err)
	}

	if err := createTable(db); err != nil {
		return nil, fmt.Errorf("failed to create table %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to database is sucefull")
	return db, nil
}

func CloseDB() {
	db.Close()
}

func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS tableOne()
	   id INTEGER PRIMARY KEY AUTOINCREMENT
	   userName TEXT NOT NULL
	   email TEXT NOT NULL
	   created_at TEXT NOT NULL);`

	_, err := db.Exec(query)
	return err
}
