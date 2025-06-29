package config

import "log/slog"

type Config struct {
	PSQL
	JWT
	Web
}

type JWT struct {
	Issuer      string `long:"jwt-issuer" env:"SERVER_JWT_ISSUER"`
	SecretKey   string `long:"jwt-secret-key" env:"SERVER_JWT_SECRET_KEY"`
	KeyLength   int    `long:"jwt-key-length" env:"SERVER_JWT_KEY_LENGTH"`
	HourExpired int    `long:"jwt-hour-expired" env:"SERVER_JWT_HOUR_EXPIRED"`
}

type Web struct {
	Host         string     `long:"web-host" env:"SERVER_WEB_HOST" default:"localhost"`
	Port         int        `long:"web-port" env:"SERVER_WEB_PORT" default:"8080"`
	ReadTimeout  int        `long:"web-read-timeout" env:"SERVER_WEB_READ_TIMEOUT"`
	WriteTimeout int        `long:"web-write-timeout" env:"SERVER_WEB_WRITE_TIMEOUT"`
	IdleTimeout  int        `long:"web-idle-timeout" env:"SERVER_WEB_IDLE_TIMEOUT"`
	Cors         string     `long:"web-cors" env:"SERVER_WEB_CORS"`
	SSLSertPath  string     `long:"web-ssl-sert-path" env:"SERVER_WEB_SSL_SERT_PATH"`
	SSLKeyPath   string     `long:"web-ssl-web" env:"SERVER_WEB_SSL_KEY_PATH"`
	CookieName   string     `long:"web-cookie-name" env:"SERVER_WEB_COOKIE_NAME"`
	SameSiteNone int        `long:"web-same-site-none" env:"SERVER_WEB_SAME_SITE_NONE"`
	LogLevel     slog.Level `long:"web-log-level" env:"SERVER_WEB_LOG_LEVEL"`
}

type PSQL struct {
	Host           string `long:"psql-host" env:"SERVER_PSQL_HOST"`
	Port           int    `long:"psql-port" env:"SERVER_PSQL_PORT" `
	UserName       string `long:"psql-user-name" env:"SERVER_PSQL_USER_NAME" `
	DBName         string `long:"psql-db-name" env:"SERVER_PSQL_DB_NAME" `
	Password       string `long:"psql-password" env:"SERVER_PSQL_PASSWORD" `
	SslMode        string `long:"psql-ssl-mode" env:"SERVER_PSQL_SSL_MODE" `
	DSN            string `long:"psql-dsn" env:"SERVER_PSQL_DSN" `
	Binary         bool   `long:"psql-binary" env:"SERVER_PSQL_BINARY"`
	MaxConnections int    `long:"psql-max-connections" env:"SERVER_PSQL_MAX_CONNECTIONS"`
	ConnectionIdle int    `long:"psql-connection-idle" env:"SERVER_PSQL_CONNECTION_IDLE"`
	MaxLimit       int    `long:"psql-max-limit" env:"SERVER_PSQL_MAX_LIMIT"`
}
