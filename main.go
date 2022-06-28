package main

import (
	"log"
	"time"

	"apoteker.id_backend/config"
	"apoteker.id_backend/router"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	// create logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Unable to create logger %s", err)
	}

	// fiber config
	fiberConf := &fiber.Config{
		AppName:      "apoteker_backend v1.0.0",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	}

	// create new fiber app
	app := fiber.New(*fiberConf)

	router.SetupRoutes(app)

	// get envs
	port := config.Config("PORT")

	// listen to port
	app.Listen(":" + port)
}
