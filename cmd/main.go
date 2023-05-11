package main

import (
	"context"
	"fmt"
	"log"

	"github.com/np-inprove/server/internal/config"
	"github.com/np-inprove/server/internal/ent"
	"github.com/np-inprove/server/internal/logger"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	var appLogger logger.AppLogger
	if cfg.App.Production {
		appLogger, err = logger.NewZapProduction()
	} else {
		appLogger, err = logger.NewZapDevelopment()
	}

	if err != nil {
		log.Fatal(err)
	}

	appLogger.Info("Hello, world!")

	client, err := ent.Open(cfg.Database.DriverName, cfg.Database.DataSourceName)
	if err != nil {
		appLogger.Fatal(fmt.Sprintf("failed opening connection to sqlite: %v", err))
	}
	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)

	if err := client.Schema.Create(context.Background()); err != nil {
		appLogger.Fatal(fmt.Sprintf("failed creating schema resources: %v", err))
	}
}
