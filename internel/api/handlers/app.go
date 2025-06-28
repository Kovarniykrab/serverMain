package handlers

import (
	"context"
	"log/slog"

	"gitlab.com/kovarniykrab/servermain/config"
	"gitlab.com/kovarniykrab/servermain/internel/service"
)

type App struct {
	cfg     *config.Config
	service *service.App
	logs    *slog.Logger
}

func New(ctx context.Context, cfg *config.Config, service *service.App, logs *slog.Logger) *App {
	app := &App{cfg: cfg, service: service, logs: logs}

	return app
}
