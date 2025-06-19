package handlers

import (
	"context"
	"github.com/rs/zerolog"
	"gitlab.com/kovarniykrab/servermain/config"
	"gitlab.com/kovarniykrab/servermain/internel/service"
)

type App struct {
	cfg     *config.Config
	service *service.App
	logs    *zerolog.Logger
}

func New(ctx context.Context, cfg *config.Config, service *service.App, logs *zerolog.Logger) *App {
	app := &App{cfg: cfg, service: service, logs: logs}

	return app
}
