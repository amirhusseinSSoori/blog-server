package bootstrap

import (
	"blog/internal/database/seeder"
	"blog/pkg/config"
	"blog/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()

}
