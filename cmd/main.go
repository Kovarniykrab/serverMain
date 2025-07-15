package main

import (
	"context"
	"database/sql"
	"embed"
	"log/slog"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jessevdk/go-flags"
	"github.com/pressly/goose/v3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uptrace/bun/driver/pgdriver"
	embedServer "gitlab.com/kovarniykrab/servermain"
	"gitlab.com/kovarniykrab/servermain/config"
	"gitlab.com/kovarniykrab/servermain/internel/api/router"
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
			ReadTimeout:       time.Second * time.Duration(app.Config.ReadTimeout),
			WriteTimeout:      time.Second * time.Duration(app.Config.Web.WriteTimeout),
			IdleTimeout:       time.Second * time.Duration(app.Config.Web.IdleTimeout),
			ReadHeaderTimeout: time.Second * time.Duration(app.Config.Web.ReadTimeout),
		}

		if err := srv.ListenAndServe(); err != nil {
			log.Error("metrics server error", "error", err)
		}
	}()

	go func() {
		if cnf.Web.SSLSertPath != "" && cnf.Web.SSLKeyPath != "" {
			if err := srv.ListenAndServeTLS(cnf.Web.SSLSertPath, cnf.Web.SSLKeyPath); err != nil {
				log.Error("server start error", "error", err)

				panic(err)
			}

			return
		}

		if err := srv.ListenAndServe(); err != nil {
			log.Error("server start error", "error", err)
			panic(err)
		}
	}()

	log.Warn("internal server started",
		"host", app.Config.Web.Host,
		"port", app.Config.Web.Port)

	c := make(chan os.Signal, 1)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(app.Config.ReadTimeout))

	if err := srv.Shutdown(ctx); err != nil {
		cancel()

		panic(err)
	}

	log.Info("server stopped")

	cancel()

	runtime.Goexit()
}

func initLogger(level *slog.Level) *slog.Logger {
	var logLevel slog.Level

	if level == nil {
		logLevel = slog.LevelInfo
	} else {
		logLevel = *level
	}
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))

	return log
} //заменить пакет логирования на slogLogger, залогировать весь код с его помощью где только возможно
//научиться дебажить в вскоде
// подтянуть конфигуранионные файлы ".env"
// развернуть приложение на сервер, ci/cd, docker
//сваггер

func initServer(ctx context.Context) (config.Config, *http.Server, *router.App, *slog.Logger) {
	conf := config.Config{}
	parser := flags.NewParser(&conf, flags.Default)
	if _, err := parser.Parse(); err != nil {
		panic(err)
	}

	log := initLogger(&conf.LogLevel)

	log.Debug("configuration initialized", "config", conf)

	repo, e := database.New(conf, log)
	if e != nil {
		log.Error("failed to initialize database", "error", e)
		panic(e)
	}

	app, err := router.New(ctx, &conf, log, repo)
	if err != nil {
		log.Error("failed to initiazile routers", "error", err)
		panic(err)
	}

	migrate(conf)

	srv := &http.Server{
		Handler:           app.GetRouter(),
		Addr:              app.Config.Web.Host + ":" + strconv.Itoa(app.Config.Web.Port),
		ReadTimeout:       time.Second * time.Duration(app.Config.Web.ReadTimeout),
		WriteTimeout:      time.Second * time.Duration(app.Config.Web.WriteTimeout),
		IdleTimeout:       time.Second * time.Duration(app.Config.Web.IdleTimeout),
		ReadHeaderTimeout: time.Second * time.Duration(app.Config.Web.ReadTimeout),
	}

	return conf, srv, app, log
}

var embedMigrations embed.FS

func migrate(cfg config.Config) {
	goose.SetBaseFS(embedServer.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.PSQL.DSN)))

	if err := goose.Up(db, "resources/store/psql/migrations"); err != nil {
		panic(err)
	}
}

// @title           Swagger report API
// @version         1.0
// @description     serverMain api server.
// @termsOfService  http://kovarniykrab.duckdns.org

// @contact.name   API Support
// @contact.url     http://kovarniykrab.duckdns.org
// @contact.email  Zaratos1999@gmmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      0.0.0.0:8080
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
