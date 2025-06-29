package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/cadyrov/goerr/v2"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"gitlab.com/kovarniykrab/servermain/config"
)

var ErrNotFound = errors.New("not found")

type Service struct {
	log    *zerolog.Logger
	sqlxDB *sqlx.DB
	db     *bun.DB
}

func New(config config.Config, log *zerolog.Logger) (*Service, goerr.IError) {
	r := Service{
		log: log,
	}

	//new bun
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.PSQL.DSN)))
	r.db = bun.NewDB(sqlDB, pgdialect.New())

	return &r, nil
}

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
		return bun.Tx{}, errors.New(fmt.Sprintf("err tx %s", err))
	}

	return tx, nil
}

func (db *Service) Ping(ctx context.Context) error {
	if err := db.db.PingContext(ctx); err != nil {
		return errors.New(fmt.Sprintf("err tx %s", err))
	}

	return nil
}
