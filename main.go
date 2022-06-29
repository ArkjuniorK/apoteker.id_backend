package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"apoteker.id_backend/config"
	router "apoteker.id_backend/routers"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
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
	app.Use(logger.New())

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
