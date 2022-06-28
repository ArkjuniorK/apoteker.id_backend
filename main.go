package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"apoteker.id_backend/config"
	"apoteker.id_backend/routers"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	// setting up
	port := config.Config("PORT")
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

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

	// register routes here
	routers.SetupRoutes(app)

	// run the server
	go func() {
		// listen to port
		err := app.Listen(":" + port)
		if err != nil {
			logger.Sugar().Errorf("Unable to start server %s", err)
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
	logger.Sugar().Info("Process terminated, gracefully shutdown the app ", sig)
	app.Shutdown()
}
