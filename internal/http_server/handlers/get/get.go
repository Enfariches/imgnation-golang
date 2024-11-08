package get

import (
	"fmt"
	resp "img/internal/lib/api/response"
	"img/internal/lib/logger/sl"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func GetImage(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http_server.handlers.get"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		uuid := chi.URLParam(r, "uuid")
		if uuid == "" {
			render.JSON(w, r, resp.Error("uuid is empty"))
			return
		}

		format := r.URL.Query().Get("extension")
		path := fmt.Sprintf("%s.%s", uuid, format)

		octet, err := os.ReadFile(filepath.Join("./upload", path))
		if err != nil {
			log.Error("Error to read file", sl.Error(err))
			render.JSON(w, r, resp.Error("Image is not found"))
			return
		}

		log.Info("Got image correctly")
		render.Data(w, r, octet) //Отправка qr-code в виде application/octet-stream
		//render.JSON(w, r, resp.OK())
	}
}
