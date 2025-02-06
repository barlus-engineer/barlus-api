package database

import "github.com/barlus-engineer/barlus-api/Internal/core/model"

func MigrateDatabase() {
	db.AutoMigrate(
		&model.User{},
	)
}