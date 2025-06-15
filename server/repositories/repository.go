package repositories

import (
	"database/sql"
	"log"
	"todoapp/models"
)

type TodoRepository struct {
	db *sql.DB
}

