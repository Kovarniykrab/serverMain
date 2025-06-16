package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

type database struct {
	*sql.DB
}

var dbInstance *database

// InitDB инициализирует подключение
func InitDB() (Database, error) {
	connStr := "user=youruser dbname=tableOne password=yourpass sslmode=disable"
	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	if err := createTable(sqlDB); err != nil {
		return nil, fmt.Errorf("error creating table: %v", err)
	}

	dbInstance = &database{sqlDB}
	log.Println("Database connection established")
	return dbInstance, nil
}

func (db *database) CloseDB() error {
	return db.DB.Close()
}

func createTable(db *sql.DB) error {
	query := `
CREATE TABLE IF NOT EXISTS tableOne (
    id SERIAL PRIMARY KEY,
    userName TEXT NOT NULL,
    email TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
)`

	_, err := db.Exec(query)
	return err
}
