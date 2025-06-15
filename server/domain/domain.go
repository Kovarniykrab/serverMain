package models

import "time"

type Model struct {
	ID        int       `json: "id"`
	UserName  string    `json: "userName"`
	Email     string    `json: "email"`
	CreatedAt time.Time `json: "created_at"`
}
