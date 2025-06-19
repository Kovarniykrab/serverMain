package config

import "github.com/rs/zerolog"

type Config struct {
	PSQL
	JWT
	Web
}

type JWT struct {
	Issuer      string `long:"jwt-issuer" env:"MAYAK_JWT_ISSUER"`
	SecretKey   string `long:"jwt-secret-key" env:"MAYAK_JWT_SECRET_KEY"`
	KeyLength   int    `long:"jwt-key-length" env:"MAYAK_JWT_KEY_LENGTH"`
	HourExpired int    `long:"jwt-hour-expired" env:"MAYAK_JWT_HOUR_EXPIRED"`
}

type Web struct {
	Host         string        `long:"web-host" env:"MAYAK_WEB_HOST" envDefault:"localhost"`
	Port         int           `long:"web-port" env:"MAYAK_WEB_PORT" envDefault:"8080"`
	ReadTimeout  int           `long:"web-read-timeout" env:"MAYAK_WEB_READ_TIMEOUT"`
	WriteTimeout int           `long:"web-write-timeout" env:"MAYAK_WEB_WRITE_TIMEOUT"`
	IdleTimeout  int           `long:"web-idle-timeout" env:"MAYAK_WEB_IDLE_TIMEOUT"`
	Cors         string        `long:"web-cors" env:"MAYAK_WEB_CORS"`
	SSLSertPath  string        `long:"web-ssl-sert-path" env:"MAYAK_WEB_SSL_SERT_PATH"`
	SSLKeyPath   string        `long:"web-ssl-web" env:"MAYAK_WEB_SSL_KEY_PATH"`
	CookieName   string        `long:"web-cookie-name" env:"MAYAK_WEB_COOKIE_NAME"`
	SameSiteNone int           `long:"web-same-site-none" env:"MAYAK_WEB_SAME_SITE_NONE"`
	LogLevel     zerolog.Level `long:"project-log-level" env:"MAYAK_PROJECT_LOG_LEVEL" required:"true"`
}

type PSQL struct {
	Host           string `long:"psql-host" env:"MAYAK_PSQL_HOST"`
	Port           int    `long:"psql-port" env:"MAYAK_PSQL_PORT"`
	UserName       string `long:"psql-user-name" env:"MAYAK_PSQL_USER_NAME"`
	DBName         string `long:"psql-db-name" env:"MAYAK_PSQL_DB_NAME"`
	Password       string `long:"psql-password" env:"MAYAK_PSQL_PASSWORD"`
	SslMode        string `long:"psql-ssl-mode" env:"MAYAK_PSQL_SSL_MODE"`
	DSN            string `long:"psql-dsn" env:"MAYAK_PSQL_DSN" required:"true"`
	Binary         bool   `long:"psql-binary" env:"MAYAK_PSQL_BINARY"`
	MaxConnections int    `long:"psql-max-connections" env:"MAYAK_PSQL_MAX_CONNECTIONS"`
	ConnectionIdle int    `long:"psql-connection-idle" env:"MAYAK_PSQL_CONNECTION_IDLE"`
	MaxLimit       int    `long:"psql-max-limit" env:"MAYAK_PSQL_MAX_LIMIT"`
}
