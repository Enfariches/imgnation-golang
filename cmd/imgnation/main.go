package main

import (
	"img/internal/config"
	"img/internal/http_server/handlers/get"
	"img/internal/http_server/handlers/save"
	mwLogger "img/internal/http_server/middleware"
	"img/internal/lib/logger/sl"
	"img/internal/logger"
	"img/internal/storage/redis"
	"img/internal/storage/s3"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.NewConfig("local")
	log := logger.SetupLogger(cfg.Env)

	db, err := s3.New() //AWS Config
	if err != nil {
		log.Error("Failed connect to S3", sl.Error(err))
	}

	redis, err := redis.New() //Redis Config
	if err != nil {
		log.Error("Failed connect to Redis", sl.Error(err))
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(mwLogger.New(log))

	r.Post("/api/save", save.SaveImage(cfg.Server.Address, log, db))
	r.Get("/api/img/{key}", get.GetImage(log, db, redis))
	// IMG -> GET UUID -> SAVE TO UPLOAD && SAVE IN BASE -> CREATE URL && CREATE QR-CODE)
	// SWITCH UPLOAD TO S3 STORAGE (AWS SDK)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/static/index.html")
	})

	r.Get("/img/{key}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/show/show.html")
	})

	r.Get("/web/static/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/web/static/", http.FileServer(http.Dir("web/static"))).ServeHTTP(w, r)
	})

	r.Get("/web/show/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/web/show/", http.FileServer(http.Dir("web/show"))).ServeHTTP(w, r)
	})

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
