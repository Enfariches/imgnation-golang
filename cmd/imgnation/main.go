package main

import (
	"img/internal/config"
	"img/internal/http_server/handlers"
	mwLogger "img/internal/http_server/middleware"
	"img/internal/lib/logger/sl"
	"img/internal/logger"
	"img/internal/storage/postgres"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.NewConfig("local")
	log := logger.SetupLogger(cfg.Env)

	storage, err := postgres.New(cfg.StorageURL)
	if err != nil {
		log.Error("Failed to init Storage", sl.Error(err))
	}
	_ = storage

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(mwLogger.New(log))

	r.Post("/api/save", handlers.SavePath(log, storage))

	log.Info("Starting server")

	server := http.Server{
		Addr:    cfg.Address,
		Handler: r,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Error("Failed to Listen and Server Server")
		os.Exit(1)
	}
	log.Info("Server stopped")

}
