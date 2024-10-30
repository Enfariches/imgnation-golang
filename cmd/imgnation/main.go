package main

import (
	"img/internal/config"
	"img/internal/logger"
	"img/internal/storage/postgres"
	"log/slog"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.NewConfig("local")
	log := logger.SetupLogger(cfg.Env)
	log.Info("Server is running", slog.String("envCfg", cfg.Env))

	storage, err := postgres.New(cfg.StorageURL)
	if err != nil {
		log.Error("Failed to init Storage", slog.String("err", err.Error()))
	}
	_ = storage
}
