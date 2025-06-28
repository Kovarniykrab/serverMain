package service

import (
	"context"

	"log/slog"

	"gitlab.com/kovarniykrab/servermain/config"
	"gitlab.com/kovarniykrab/servermain/internel/domain"
)

type App struct {
	Database Database
	cfg      *config.Config
	logger   *slog.Logger
	ctx      context.Context
	db       *Database
}

type Database interface {
	UsersSearch(ctx context.Context, form domain.UserSearchForm) (domain.Users, int, error)
	GetUserByID(ctx context.Context, id int) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
}

func New(ctx context.Context, cfg *config.Config, logger *slog.Logger, Database Database) *App {

	return &App{
		cfg:      cfg,
		logger:   logger,
		Database: Database,
		ctx:      ctx,
	}
}
