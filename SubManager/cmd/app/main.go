package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lashkapashka/SubManager/docs"
	"github.com/lashkapashka/SubManager/internal/config"
	subscriptioncreate "github.com/lashkapashka/SubManager/internal/http-server/handlers/subscriptionCreate"
	subscriptiondelete "github.com/lashkapashka/SubManager/internal/http-server/handlers/subscriptionDelete"
	subscriptionget "github.com/lashkapashka/SubManager/internal/http-server/handlers/subscriptionGet"
	subscriptiontotal "github.com/lashkapashka/SubManager/internal/http-server/handlers/subscriptionTotalPrice"
	subscriptionupdate "github.com/lashkapashka/SubManager/internal/http-server/handlers/subscriptionUpdate"
	"github.com/swaggo/http-swagger"

	"github.com/lashkapashka/SubManager/internal/service"
	"github.com/lashkapashka/SubManager/internal/storage/postgresql"
)

const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)

// @title           Subscription Manager API
// @version         1.0
// @description     API для работы с подписками пользователей.
// @host            localhost:8080
// @BasePath        /

func main() {
	cfg := config.MustLoad()
	logger := setupLogger(cfg.Env)

	logger.Info(
		"starting subManager",
		slog.String("env", cfg.Env),
		slog.String("version", "@1.0.1"),
	)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	storage := postgresql.New(cfg.StoragePath, logger)

	service := service.New(storage, logger)

	router.Group(func(r chi.Router) {
		r.Post("/subscriptions", subscriptioncreate.CreateSubscription(service, logger))
		r.Get("/subscriptions/total", subscriptiontotal.TotalPrice(service, logger))
		r.Get("/subscriptions", subscriptionget.Get(service, logger))
		r.Patch("/subscriptions", subscriptionupdate.Update(service, logger))
		r.Delete("/subscriptions", subscriptiondelete.Delete(service, logger))
	})

	logger.Info("starting server", slog.String("address", cfg.Address))
	
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr: cfg.Address,
		Handler: router,
		ReadTimeout: cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout: cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error("failed to stop server")
		}
	}()
	
	logger.Info("server started")
	
	<-done
	logger.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("failed to stop server")
		return
	}

	logger.Info("server stopped")
}	

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log  = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}