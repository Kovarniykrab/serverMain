package router

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gorilla/mux"
	"gitlab.com/kovarniykrab/servermain/config"
	"gitlab.com/kovarniykrab/servermain/internel/database"
	"gitlab.com/kovarniykrab/servermain/internel/service"

	"gitlab.com/kovarniykrab/servermain/internel/api/handlers"
)

type App struct {
	handlers.App
	Config  *config.Config
	logger  *slog.Logger
	Service *service.App
}

func New(ctx context.Context, cfg *config.Config, logger *slog.Logger) (*App, error) {
	var repo service.Database
	fmt.Println(cfg.PSQL.DSN)
	repo, e := database.New(*cfg, logger)
	if e != nil {
		logger.Error("failed to initialize database", "error", e)
		panic(e)
	}

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
