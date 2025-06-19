package router

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"gitlab.com/kovarniykrab/servermain/config"
	"gitlab.com/kovarniykrab/servermain/internel/service"

	"gitlab.com/kovarniykrab/servermain/internel/api/handlers"
)

type App struct {
	handlers.App
	Config  *config.Config
	logger  *zerolog.Logger
	Service *service.App
}

func New(ctx context.Context, cfg *config.Config, logger *zerolog.Logger, repo service.Database) (*App, error) {

	srv := service.New(ctx, cfg, logger, repo)
	hand := handlers.New(ctx, cfg, srv, logger)

	return &App{
		App:     *hand,
		Config:  cfg,
		logger:  logger,
		Service: srv,
	}, nil
}

func (app *App) GetRouter() *mux.Router {

	routes := mux.NewRouter()
	api := routes.PathPrefix("/api").Subrouter()
	setupRouter := api.PathPrefix("/user").Subrouter()
	app.SetupUserRoutes(setupRouter)

	return routes
}
