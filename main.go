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
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

var db *gorm.DB

// fiber config
var fiberConf = &fiber.Config{
	AppName:      "apoteker_backend v1.0.0",
	ReadTimeout:  30 * time.Second,
	WriteTimeout: 30 * time.Second,
	IdleTimeout:  10 * time.Second,
	JSONEncoder:  json.Marshal,
	JSONDecoder:  json.Unmarshal,
}

var loggerConf = &logger.Config{
	Format:     "${time}\t${method}\b${path}\t${status}\n",
	TimeZone:   "Asia/Makassar",
	TimeFormat: "2 Jan 2006 15:04:05",
}

func main() {
	// get all envs
	port := config.Config("PORT")

	log := InitLogger()
	defer log.Sync()

	// create new fiber app
	// setup middlewares and routers
	app := fiber.New(*fiberConf)
	app.Use(logger.New(*loggerConf))

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

	// connect database
	// get the sql from db
	// it would be used to close the connection
	db = database.ConnectDB(log)
	sql, err := db.DB()
	if err != nil {
		log.Sugar().Fatal("An error has occured\n", err)
	}

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

	// kill other tasks
	log.Sugar().Info("Running cleanup tasks...")
	sql.Close()
	log.Sugar().Info("App successfully shutdown")
}

// function to initate logger
// with default configuration
// @todo:
// - use env to determine different config
func InitLogger() *zap.Logger {
	// set logger configuration
	conf := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Encoding:         "console",
		Development:      false,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "M",
			LevelKey:      "L",
			TimeKey:       "T",
			StacktraceKey: "S",
			CallerKey:     "C",
			NameKey:       "N",

			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("2 Jan 2006 15:04:05"),
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	// build configuration
	log, err := conf.Build()
	if err != nil {
		panic(err)
	}

	// return the log
	return log
}
