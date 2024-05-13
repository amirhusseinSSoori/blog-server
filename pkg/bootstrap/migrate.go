package bootstrap

import (
	"blog/internal/database/migration"
	"blog/pkg/config"
	"blog/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()

}
