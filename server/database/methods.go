package database

import (
	"database/sql"

	"gitlab.com/kovarniykrab/servermain/server/domain"
)

type database struct {
	*sql.DB
}

func (db *database) GetAllUsers() ([]domain.Model, error) {
	rows, err := db.Query("SELECT id, userName, email, craeted_at FROM tableOne")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.Model

	for rows.Next() {
		var u domain.Model
		if err := rows.Scan(&u.ID, &u.UserName, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (db *database) GetUserByID(id int) (*domain.Model, error) {
	var u domain.Model
	err := db.QueryRow("SELECT id, userName, email, created_at FROM tableOne WHERE 	id $1", id).Scan(&u.ID, &u.UserName, &u.Email, &u.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (db *database) CreateUser(user domain.Model) (int, error) {
	res, err := db.Exec(
		"INSERT INTO tableOne (id, userName, email, created_at) VALUES ($1, $2, $3, $4)",
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

func (db *database) UpdateUser(id int, user domain.Model) (*domain.Model, error) {
	var updatedUser domain.Model
	err := db.QueryRow(
		"UPDATE tableOne SET userName = $1, email = $2 WHERE id = $3 RETURNING id, userName, email, created_at",
		user.UserName, user.Email, id,
	).Scan(&updatedUser.ID, &updatedUser.UserName, &updatedUser.Email, &updatedUser.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

func (db *database) DeleteUser(id int) error {
	_, err := db.Exec("DELETE FROM tableONe WHERE id = $1", id)
	return err
}
