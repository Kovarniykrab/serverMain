package main

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"gitlab.com/kovarniykrab/servermain/config"
	routers "gitlab.com/kovarniykrab/servermain/internel/api/router"
	"gitlab.com/kovarniykrab/servermain/internel/database"
)

func main() {
	ctx := context.Background()

	cnf, srv, app, log := initServer(ctx)

	go func() {
		rt := mux.NewRouter()
		rt.Handle("/metrics", promhttp.Handler())

		srv := &http.Server{
			Handler:           rt,
			Addr:              ":8080",
			ReadTimeout:       time.Second * time.Duration(app.Config.Web.ReadTimeout),
			WriteTimeout:      time.Second * time.Duration(app.Config.Web.WriteTimeout),
			IdleTimeout:       time.Second * time.Duration(app.Config.Web.IdleTimeout),
			ReadHeaderTimeout: time.Second * time.Duration(app.Config.Web.ReadTimeout),
		}

		if err := srv.ListenAndServe(); err != nil {
			log.Err(err).Msg("metrics")
		}
	}()

	go func() {
		if cnf.Web.SSLSertPath != "" && cnf.Web.SSLKeyPath != "" {
			if err := srv.ListenAndServeTLS(cnf.Web.SSLSertPath, cnf.Web.SSLKeyPath); err != nil {
				log.Err(err).Msg("start wrong")

				panic(err)
			}

			return
		}

		if err := srv.ListenAndServe(); err != nil {
			log.Err(err).Msg("start wrong")

			panic(err)
		}
	}()

	log.Warn().Str("host", app.Config.Web.Host).
		Int("port", app.Config.Web.Port).
		Msg("internel start")

	c := make(chan os.Signal, 1)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(app.Config.ReadTimeout))

	if err := srv.Shutdown(ctx); err != nil {
		cancel()

		panic(err)
	}

	log.Info().Msg("service stop")

	cancel()

	runtime.Goexit()
}

func initLogger(conf config.Config) zerolog.Logger {
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	log.WithLevel(conf.LogLevel)

	return log
}

func initServer(ctx context.Context) (config.Config, *http.Server, *routers.App, *zerolog.Logger) {
	conf := config.Config{}

	parser := flags.NewParser(&conf, flags.Default)
	if _, err := parser.Parse(); err != nil {
		fmt.Printf("error parse env: %s\n", err.Error())
		os.Exit(1)
	}
	log := initLogger(conf)

	repo, e := database.New(conf, &log)
	if e != nil {
		panic(e)
	}

	app, err := routers.New(ctx, &conf, &log, repo)
	if err != nil {
		panic(e)
	}

	//migrate(db)

	srv := &http.Server{
		Handler:           app.GetRouter(),
		Addr:              app.Config.Web.Host + ":" + strconv.Itoa(app.Config.Web.Port),
		ReadTimeout:       time.Second * time.Duration(app.Config.Web.ReadTimeout),
		WriteTimeout:      time.Second * time.Duration(app.Config.Web.WriteTimeout),
		IdleTimeout:       time.Second * time.Duration(app.Config.Web.IdleTimeout),
		ReadHeaderTimeout: time.Second * time.Duration(app.Config.Web.ReadTimeout),
	}

	return conf, srv, app, &log
}

var embedMigrations embed.FS

func migrate(db *sqlx.DB) {
	s := embedMigrations

	goose.SetBaseFS(s)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db.DB, "resources/store/psql/migrations"); err != nil {
		panic(err)
	}
}
