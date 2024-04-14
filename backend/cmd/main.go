package main

import (
	"log/slog"
	"net/http"
	"os"
	AI "smart-beyim/internal/AI"
	"smart-beyim/internal/config"
	"smart-beyim/internal/handlers"
	"smart-beyim/internal/repo"
	"smart-beyim/internal/service"
	"smart-beyim/lib/logger/sl"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// init config
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("start-up url_shotener", slog.String("env", cfg.Env))
	log.Debug("debug message are enabled")

	storage, err := repo.NewDB(cfg.StoragePath)

	// if err = storage.SeedDatabase(); err != nil {
	// 	log.Error("failed init database", sl.Err(err))
	// 	os.Exit(1)
	// }

	if err != nil {
		log.Error("failed init database", sl.Err(err))
		os.Exit(1)
	}
	log.Info("database is started")

	ai := AI.New(cfg.ApiKey)
	log.Info("ai is started")

	s := service.NewService(storage, ai)
	log.Info("starting servic")

	h := handlers.NewHandler(s, log)

	log.Info("starting server", slog.String("addres", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      h.Router(),
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to run server", err)
	}
	log.Error("stopped server")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
