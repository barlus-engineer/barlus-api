package database

import (
	"github.com/barlus-engineer/barlus-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func PostgresConnect() error {
	var (
		cfg = config.GetConfig()
		dns = cfg.Database.PostgresURL
	)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return err
	}
	
	dbClient = db

	return nil
}

func GetDatabase() *gorm.DB {
	return dbClient
}