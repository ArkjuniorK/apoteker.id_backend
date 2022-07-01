package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ArkjuniorK/apoteker.id_backend/config"
	"github.com/ArkjuniorK/apoteker.id_backend/database"
	"github.com/ArkjuniorK/apoteker.id_backend/router"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	log *zap.Logger
	db  *gorm.DB
)

// fiber config
var fiberConf = &fiber.Config{
	AppName:      "apoteker_backend v1.0.0",
	ReadTimeout:  30 * time.Second,
	WriteTimeout: 30 * time.Second,
	IdleTimeout:  10 * time.Second,
	JSONEncoder:  json.Marshal,
	JSONDecoder:  json.Unmarshal,
}

func main() {
	// get all envs
	port := config.Config("PORT")

	// set logger, database connection
	// based on given env
	log, _ = zap.NewProduction()
	defer log.Sync()

	// connect database
	// get the sql from db
	// it would be used to close the connection
	db = database.ConnectDB(log)
	sql, err := db.DB()
	if err != nil {
		log.Sugar().Fatal("An error has occured\n", err)
	}

	// create new fiber app
	// setup middlewares and routers
	app := fiber.New(*fiberConf)
	app.Use(logger.New(logger.Config{
		Format:   "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeZone: "Asia/Makassar",
	}))

	// setup router here
	router.SetupRouter(app, log)

	// run the server
	// and listen to port
	go func() {
		err := app.Listen(":" + port)
		if err != nil {
			log.Sugar().Panic("Unable to start server ", err)
		}
	}()

	// graceful shutdown mechanism
	// first create channel to check signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	// second notify the signal
	// and print info then shutdown
	s := <-c
	log.Sugar().Info("Gracefully shutting down...\n", s)
	_ = app.Shutdown()

	log.Sugar().Info("Running cleanup tasks...")

	// kill other tasks
	sql.Close()
	log.Sugar().Info("App successfully shutdown")
}
