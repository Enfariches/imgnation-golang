package get

import (
	resp "img/internal/lib/api/response"
	"img/internal/lib/logger/sl"
	"img/internal/storage"
	"img/internal/storage/redis"
	"img/internal/storage/s3"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func GetImage(log *slog.Logger, db *s3.StorageS3, redis *redis.StorageRedis) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http_server.handlers.get"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		key := chi.URLParam(r, "key")
		if key == "" {
			render.JSON(w, r, resp.Error("key is empty"))
			return
		}

		octet, err := redis.CacheGet(key) //Redis Func
		if err != nil {
			if err != storage.ErrRedisNotFoundOctet {
				log.Info("Image not found in Redis")
			}

			octet, err := db.Get(key)
			if err != nil {
				log.Error("Failed to get octet from S3", sl.Error(err))
				render.JSON(w, r, resp.Error("Error to get image"))
				return
			}

			err = redis.CacheSave(octet, key)
			if err != nil {
				log.Error("Failed to save in Redis", sl.Error(err))
			}

			if err != nil {
				log.Error("Error to read file", sl.Error(err))
				return
			}
		}

		log.Info("Got image correctly")
		render.Data(w, r, octet) //Отправка qr-code в виде application/octet-stream
	}
}
