package database

import (
	"database/sql"
	"time"
)

type database struct {
	*sql.DB
}

type User struct {
	ID        int       `json: "id"`
	UserName  string    `json: "userName"`
	Email     string    `json: "email"`
	CreatedAt time.Time `json: "created_at"`
}

func (db *database) GetAllUsers() ([]User, error) {
	rows, err := db.Query("SELECT id, userName, email, craeted_at FROM tableOne")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.UserName, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (db *database) GetUserByID(id int) (*User, error) {
	var u User
	err := db.QueryRow("SELECT id, userName, email, created_at FROM tableOne WHERE 	id ?", id).Scan(&u.ID, &u.UserName, &u.Email, &u.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (db *database) CreateUser(user User) (int, error) {
	res, err := db.Exec(
		"INSERT INTO tableOne (id, userName, email, created_at) VALUES (?, ?, ?, ?)",
		user.ID, user.UserName, user.Email, user.CreatedAt,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}

func (db *database) UpdateUser(user User) error {
	_, err := db.Exec(
		"UPDATE tableOne SET id = ?, userName = ?, email = ?, created_at = ?",
	)
	return err
}

func (db *database) DeleteUser(id int) error {
	_, err := db.Exec("DELETE FROM tableONe WHERE id = ?", id)
	return err
}
