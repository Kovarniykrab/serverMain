package servermain

import (
	"embed"
)

//go:embed resources/migrations/psql/*.sql
var EmbedMigrations embed.FS
