package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"apoteker.id_backend/config"
	"apoteker.id_backend/database/connect"
	router "apoteker.id_backend/router"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	//"gorm.io/gorm/logger"
)

func main() {
	// setting up
	port := config.Config("PORT")
	zlog, _ := zap.NewDevelopment()
	defer zlog.Sync()

	// fiber config
	fiberConf := &fiber.Config{
		AppName:      "apoteker_backend v1.0.0",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	}

	// create new fiber app
	app := fiber.New(*fiberConf)

	// middlewares
	//app.Use(logger.New())

	// connect to database
	connect.Connect()

	// routes
	router.SetupRouter(app, zlog)

	// run the server
	// and listen to port
	go func() {
		err := app.Listen(":" + port)
		if err != nil {
			zlog.Sugar().Panic("Unable to start server ", err)
		}
	}()

	// graceful shutdown mechanism
	// first create channel to check signal
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	signal.Notify(channel, syscall.SIGTERM)

	// second notify the signal
	// and print info then shutdown
	sig := <-channel
	zlog.Sugar().Info("Process terminated, gracefully shutdown the app ", sig)
	app.Shutdown()
}
