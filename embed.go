package embedServer

import (
	"embed"
)

//go:embed resources/store/psql/migrations/*.sql
var EmbedMigrations embed.FS
