package config

import (
	"time"
)

type Config struct {
	PSQL
	JWT
	Web
	S3
	Redis
	Rabbit
	Dadata
}

type JWT struct {
	Issuer      string `long:"jwt-issuer" env:"MAYAK_JWT_ISSUER"`
	SecretKey   string `long:"jwt-secret-key" env:"MAYAK_JWT_SECRET_KEY"`
	KeyLength   int    `long:"jwt-key-length" env:"MAYAK_JWT_KEY_LENGTH"`
	HourExpired int    `long:"jwt-hour-expired" env:"MAYAK_JWT_HOUR_EXPIRED"`
}

type Web struct {
	Host         string `long:"web-host" env:"MAYAK_WEB_HOST"`
	Port         int    `long:"web-port" env:"MAYAK_WEB_PORT"`
	ReadTimeout  int    `long:"web-read-timeout" env:"MAYAK_WEB_READ_TIMEOUT"`
	WriteTimeout int    `long:"web-write-timeout" env:"MAYAK_WEB_WRITE_TIMEOUT"`
	IdleTimeout  int    `long:"web-idle-timeout" env:"MAYAK_WEB_IDLE_TIMEOUT"`
	Cors         string `long:"web-cors" env:"MAYAK_WEB_CORS"`
	SSLSertPath  string `long:"web-ssl-sert-path" env:"MAYAK_WEB_SSL_SERT_PATH"`
	SSLKeyPath   string `long:"web-ssl-web" env:"MAYAK_WEB_SSL_KEY_PATH"`
	CookieName   string `long:"web-cookie-name" env:"MAYAK_WEB_COOKIE_NAME"`
	SameSiteNone int    `long:"web-same-site-none" env:"MAYAK_WEB_SAME_SITE_NONE"`
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

type S3 struct {
	Endpoint  string `long:"s3-endpoint" env:"MAYAK_S3_ENDPOINT"`
	AccessKey string `long:"s3-access-key" env:"MAYAK_S3_ACCESS_KEY"`
	SecretKey string `long:"s3-secret-key" env:"MAYAK_S3_SECRET_KEY"`
	Bucket    string `long:"s3-bucket" env:"MAYAK_S3_BUCKET"`
	Region    string `long:"s3-region" env:"MAYAK_S3_REGION"`
}

type Redis struct {
	RedisHost     string        `long:"redis-host" env:"MAYAK_REDIS_HOST" description:"хост"`
	RedisPort     string        `long:"redis-port" env:"MAYAK_REDIS_PORT" description:"порт"`
	RedisPassword string        `long:"redis-password" env:"MAYAK_REDIS_PASSWORD" description:"пароль"`
	RedisTTL      time.Duration `long:"redis-ttl" env:"MAYAK_REDIS_TTL"`
	RedisDB       int           `long:"redis-db" env:"MAYAK_REDIS_DB"`
	RedisKey      string        `long:"redis-key" env:"MAYAK_REDIS_KEY"`
	RedisUserName string        `long:"redis-username" env:"MAYAK_REDIS_USERNAME"`
}

type Rabbit struct {
	Host      string `long:"notifier-host" env:"MAYAK_RABBIT_HOST"`
	Port      int    `long:"notifier-port" env:"MAYAK_RABBIT_PORT"`
	UserName  string `long:"notifier-user-name" env:"MAYAK_RABBIT_USER_NAME"`
	Password  string `long:"notifier-password" env:"MAYAK_RABBIT_PASSWORD"`
	NameQueue string `long:"notifier-name-queue" env:"MAYAK_RABBIT_NAME_QUEUE"`
	VHost     string `long:"notifier-vhost" env:"MAYAK_RABBIT_VHOST"`
}

type Dadata struct {
	APIKey    string `long:"dadata-api-key" env:"MAYAK_DADATA_API_KEY" required:"true"`
	APIURI    string `long:"dadata-api-uri" env:"MAYAK_DADATA_API_URI" required:"true"`
	DebugMode bool   `long:"dadata-debug-mode" env:"MAYAK_DADATA_DEBUG_MODE"`
}
