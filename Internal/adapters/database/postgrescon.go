package database

import (
	"github.com/barlus-engineer/barlus-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	dblogger "gorm.io/gorm/logger"
)

var db *gorm.DB

func PostgresConnect() error {
	var (
		err error
		cfg = config.GetConfig()
		dns = cfg.Database.PostgresURL
		
		dbLoggerMode dblogger.LogLevel
	)
	
	if cfg.Release {
		dbLoggerMode = dblogger.Silent
	} else {
		dbLoggerMode = dblogger.Info
	}

	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: dblogger.Default.LogMode(dbLoggerMode),
	})
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}