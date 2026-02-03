package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/cadyrov/goerr/v2"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"gitlab.com/kovarniykrab/servermain/config"
)

var ErrNotFound = errors.New("not found")

type Service struct {
	log *slog.Logger
	db  *bun.DB
}

func connectSQLDB(dsn string) (*sql.DB, error) {
	fmt.Println(9)
	connector := pgdriver.NewConnector(pgdriver.WithDSN(dsn))
	fmt.Println(7)
	sqldb := sql.OpenDB(connector)
	return sqldb, nil
}
func New(config config.Config, log *slog.Logger) (*Service, goerr.IError) {
	fmt.Println(6)
	sqldb, err := connectSQLDB(config.DSN)
	fmt.Println(7)
	if err != nil {
		log.Debug("Database not ready...", "err", err)
		return nil, nil
	}
	if err := sqldb.PingContext(context.Background()); err != nil {
		log.Debug("Ping failed...", "err", err)
		return nil, nil
	}
	bunDB := bun.NewDB(sqldb, pgdialect.New())
	db := &Service{
		db:  bunDB,
		log: log,
	}

	return db, nil
}

//nolint:goerr113
