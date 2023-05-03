package main

import (
	"log"

	"github.com/np-inprove/server/internal/config"
	"github.com/np-inprove/server/internal/logger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	var appLogger logger.AppLogger
	if cfg.App.Environment == "Production" {
		appLogger, err = logger.NewZapProduction()
	} else {
		appLogger, err = logger.NewZapDevelopment()
	}

	if err != nil {
		log.Fatal(err)
	}

	appLogger.Info("Hello, world!")
}
