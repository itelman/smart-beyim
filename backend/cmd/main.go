package main

import (
	"log/slog"
	"os"
	"smart-beyim/internal/config"
	"smart-beyim/internal/repo"
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

	if err != nil {
		log.Error("failed init database", sl.Err(err))
		os.Exit(1)
	}
	_ = storage
	log.Info("database is started")

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
