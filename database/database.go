package database

import (
	"apoteker.id_backend/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error

func ConnectDB(env string, log *zap.Logger) *gorm.DB {
	dsn := config.Config("DSN")

	// if env is development connect to mysql
	if env == "development" {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Sugar().Fatalf("Unable to connect to database. \n", err)
		}

		log.Info("Database connected")
		return db
	}

	// otherwise connect to postgres
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Sugar().Fatalf("Unable to connect to database. \n", err)
	}

	log.Info("Database connected")
	return db
}
