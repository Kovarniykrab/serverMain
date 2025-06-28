package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/cadyrov/goerr/v2"
	"github.com/jmoiron/sqlx"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"gitlab.com/kovarniykrab/servermain/config"
)

var ErrNotFound = errors.New("not found")

type Service struct {
	log    *slog.Logger
	sqlxDB *sqlx.DB
	db     *bun.DB
}

func New(config config.Config, log *slog.Logger) (*Service, goerr.IError) {
	r := Service{
		log: log,
	}

	//new bun
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.PSQL.DSN)))
	r.db = bun.NewDB(sqlDB, pgdialect.New())

	return &r, nil
}

//nolint:goerr113

func (db *Service) BeginSQLX(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, goerr.IError) {
	tx, e := db.sqlxDB.BeginTxx(ctx, opts)
	if e != nil {
		return nil, goerr.Internal(e)
	}

	return tx, nil
}

func (db *Service) TX(ctx context.Context) (bun.Tx, error) {
	tx, err := db.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return bun.Tx{}, fmt.Errorf("some error: %v", err)
	}

	return tx, nil
}
