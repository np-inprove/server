package main

import (
	"context"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/np-inprove/server/internal/httphandler"
	"github.com/np-inprove/server/internal/usecase"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/np-inprove/server/internal/config"
	"github.com/np-inprove/server/internal/ent"
	"github.com/np-inprove/server/internal/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to create app config: %v", err)
	}

	var appLogger logger.AppLogger
	if cfg.App.Production {
		appLogger, err = logger.NewZapProduction()
	} else {
		appLogger, err = logger.NewZapDevelopment()
	}

	if err != nil {
		log.Fatalf("failed to create app logger: %v", err)
	}

	appLogger.Info("initializing server: hello, world!")
	appLogger.Info("opening connection",
		logger.String("area", "database"),
		logger.String("driverName", cfg.Database.DriverName),
	)

	client, err := ent.Open(cfg.Database.DriverName, cfg.Database.DataSourceName)
	if err != nil {
		appLogger.Fatal("failed opening connection",
			logger.String("err", err.Error()),
			logger.String("area", "database"),
		)
	}

	appLogger.Info("successfully opened connection",
		logger.String("area", "database"),
		logger.String("driverName", cfg.Database.DriverName),
	)

	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)

	if cfg.Database.AutoMigration {
		appLogger.Info("running auto migration",
			logger.String("area", "database"),
			logger.String("driverName", cfg.Database.DriverName),
		)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := client.Schema.Create(ctx); err != nil {
			appLogger.Fatal("failed creating schema resources",
				logger.String("err", err.Error()),
				logger.String("area", "database"),
				logger.String("driverName", cfg.Database.DriverName),
			)
		}

		appLogger.Info("completed auto migration",
			logger.String("area", "database"),
			logger.String("driverName", cfg.Database.DriverName),
		)
	}

	appLogger.Info("creating http server",
		logger.String("area", "http"),
	)

	r := chi.NewRouter()
	loggerMiddleware := logger.NewMiddleware(appLogger)

	//tokenAuth := jwtauth.New(cfg.App.JWTAlgorithm, cfg.App.JWTSignKey, cfg.App.JWTVerifyKey)

	r.Use(cors.Handler(cors.Options{
		AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(loggerMiddleware.Request)

	auc := usecase.NewAuth()
	authHandler := httphandler.NewAuth(auc)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		render.DefaultResponder(w, r, render.M{"ok": true})
	})
	r.Mount("/auth", authHandler)

	server := http.Server{Addr: cfg.HTTP.Addr, Handler: r}

	appLogger.Info("created http server",
		logger.String("area", "http"),
		logger.String("addr", cfg.HTTP.Addr),
	)

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		err = server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			appLogger.Fatal("failed to listen",
				logger.String("err", err.Error()),
				logger.String("area", "http"),
				logger.String("addr", cfg.HTTP.Addr),
			)
		}
		wg.Done()
	}()

	go func() {
		i := <-interrupt

		appLogger.Info("interrupt detected and shutting down gracefully",
			logger.String("signal", i.String()),
		)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		appLogger.Info("shutting down http server", logger.String("area", "http"))

		err = server.Shutdown(ctx)
		if err != nil {
			appLogger.Fatal("failed to shut down http server gracefully",
				logger.String("err", err.Error()),
				logger.String("area", "http"),
			)
		}

		appLogger.Info("successfully shut down http server", logger.String("area", "http"))
		appLogger.Info("closing connection with database", logger.String("area", "database"))

		err = client.Close()
		if err != nil {
			appLogger.Fatal("failed to close database connection",
				logger.String("err", err.Error()),
				logger.String("area", "database"),
			)
		}

		appLogger.Info("closed connection with database",
			logger.String("area", "database"),
		)
	}()

	wg.Wait()

	appLogger.Info("completed graceful shutdown, bye!")

	//r.Use(jwtauth.Verifier(tokenAuth))
	//r.Use(jwtauth.Authenticator)

}
