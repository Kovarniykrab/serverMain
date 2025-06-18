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

	"server/config"

	"github.com/gorilla/mux"
	"github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	goose "github.com/pressly/goose/v3"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/rs/zerolog"
	routers "gitlab.com/kovarniykrab/servermain/server/api/router"
)

func main() {
	cnf, srv, app, log := initServer()
	ctx := context.Background()

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

	go func() {
		app.Service.WithSyncFuncs(
			app.Service.SyncNomenclature,            // moved
			app.Service.SyncPayments,                // moved to s3
			app.Service.SyncOrder,                   // moved to s3
			app.Service.SyncColors,                  // moved to s3
			app.Service.SyncClients,                 // moved
			app.Service.SyncBooking,                 // moved to s3
			app.Service.SyncStatusOrderInProduction, // moved to s3
			app.Service.SyncStatusOrderInStock,      // moved to s3
			app.Service.SyncStatusOrderInWork,       // moved to s3
			app.Service.SyncStatusOrderInShipped,    // moved to s3
			app.Service.SyncCompany,                 // moved to s3
			app.Service.SyncUserLinkCompany,
		)

		log.Warn().Msg("sync func served")

		// app.Service.RunSyncLocal(ctx, app.Config.FilesTTL)
		app.Service.RunSync(ctx, app.Config.FilesTTL)
	}()

	log.Warn().Str("host", app.Config.Web.Host).
		Int("port", app.Config.Web.Port).
		Msg("server start")

	c := make(chan os.Signal, 1)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(app.Config.ReadTimeout))

	if err := srv.Shutdown(ctx); err != nil {
		cancel()

		panic(err)
	}

	log.Info().Msg("server stop")

	cancel()

	runtime.Goexit()
}

func initLogger(conf config.Config) zerolog.Logger {
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	log.WithLevel(conf.LogLevel)

	return log
}

func initServer() (config.Config, *http.Server, *routers.App, *zerolog.Logger) {
	conf := config.Config{}

	parser := flags.NewParser(&conf, flags.Default)
	if _, err := parser.Parse(); err != nil {
		fmt.Printf("error parse env: %s\n", err.Error())
		os.Exit(1)
	}

	log := initLogger(conf)

	repo, db, e := store.New(conf, &log)
	if e != nil {
		panic(e)
	}

	app := routers.New(conf, &log, service.TypeAuth, repo)

	migrate(db)

	go func() {
		app.Service.SendMail()
	}()

	srv := &http.Server{
		Handler:           app.GetRoutes(),
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
