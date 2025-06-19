package domain

import "time"

type User struct {
	ID        int        `bun:"id,pk,autoincrement" json:"id"`
	UserName  string     `bun:"user_name" json:"user_name"`
	Email     string     `bun:"email" json:"email"`
	CreatedAt *time.Time `bun:"created_at" json:"created_at"`
}

type Users []User

type UserSearchForm struct {
	UserName string `json:"user_name"`
	SearchForm
}
