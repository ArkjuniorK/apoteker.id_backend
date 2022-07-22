package database

import (
	"fmt"

	"github.com/ArkjuniorK/apoteker.id_backend/config"
	"github.com/ArkjuniorK/apoteker.id_backend/internal/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB(log *zap.Logger) *gorm.DB {
	appEnv := config.Config("APP_ENV")
	dbName := config.Config("DB_NAME")
	dbUser := config.Config("DB_USER")
	dbPass := config.Config("DB_PASS")
	dbHost := config.Config("DB_HOST")
	dbPort := config.Config("DB_PORT")
	// dbSSLMode := config.Config("DB_SSLMODE")
	dbURI := config.Config("DATABASE_URL")

	// if env is equal to development
	// connect db to mysql
	if appEnv == "development" { //use mysql
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		// check error
		if err != nil {
			log.Sugar().Fatalf("Unalbe to connect MySQL, error: %s", err)
		}

		log.Sugar().Infof("Connected to MySQL")
	} else {
		// otherwise use Postgres
		// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPass, dbName, dbSSLMode)
		DB, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		log.Sugar().Info("Connected to Postgres")
	}

	DB.AutoMigrate(&model.Apoteker{}, &model.Apotek{}, &model.Pegawai{})
	return DB
}
