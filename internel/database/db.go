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

func New(config config.Config, log *zerolog.Logger) (*Service, *sqlx.DB, goerr.IError) {
	r := Service{
		log: log,
	}

	url, e := ConnectionURL(config.PSQL)
	if e != nil {
		return nil, nil, e
	}

	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		return nil, nil, goerr.Internal(err)
	}

	r.sqlxDB = db

	//new bun
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.PSQL.DSN)))
	r.db = bun.NewDB(sqlDB, pgdialect.New())

	if err := r.Ping(context.Background()); err != nil {
		return nil, nil, goerr.Internal(err)
	}

	return &r, db, nil
}

//nolint:goerr113
func ConnectionURL(c config.PSQL) (url string, e goerr.IError) {
	url = "host=%s port=%d user=%s password=%s dbname=%s"

	if c.Host == "" || c.Port == 0 || c.UserName == "" || c.DBName == "" || c.Password == "" {
		err := fmt.Errorf("config isn't full "+url, c.Host,
			c.Port, c.UserName, c.Password, c.DBName)
		e = goerr.BadRequest(err)

		return
	}

	if c.SslMode != "" {
		url += " sslmode=" + c.SslMode
	}

	if c.Binary {
		url += " binary_parameters=yes"
	}

	url = fmt.Sprintf(url, c.Host, c.Port, c.UserName, c.Password, c.DBName)

	return
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
