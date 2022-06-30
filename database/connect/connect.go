package connect

import (
	"fmt"

	"apoteker.id_backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	dbName := config.Config("DB_NAME")
	dbUser := config.Config("DB_USER")
	dbPass := config.Config("DB_PASS")
	dbHost := config.Config("DB_HOST")
	dbPort := config.Config("DB_PORT")
	dbSSLMode := config.Config("DB_SSLMODE")

	if config.Config("APP_ENV") == "development" { //use mysql
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		fmt.Println("mysql")
		if err != nil {
			panic(err)
		}
	} else if config.Config("APP_ENV") == "production" { //use postgres
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPass, dbName, dbSSLMode)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		fmt.Println("postgres")
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Connection Opened to Database")
}
