package database

import "gitlab.com/kovarniykrab/servermain/server/domain"

type Database interface {
	GetAllUsers() ([]domain.Model, error)
	GetUserByID(id int) (*domain.Model, error)
	CreateUser(user domain.Model) (int, error)
	UpdateUser(id int, user domain.Model) (*domain.Model, error)
	DeleteUser(id int) error
	CloseDB() error
}
