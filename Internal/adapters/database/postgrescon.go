package database

import (
	"github.com/barlus-engineer/barlus-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func PostgresConnect() error {
	var (
		err error
		cfg = config.GetConfig()
		dns = cfg.Database.PostgresURL
	)
	
	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}