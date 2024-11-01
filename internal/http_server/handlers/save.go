package handlers

import (
	"fmt"
	"img/internal/lib/logger/sl"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5/middleware"
)

type Saver interface {
	SavePath(path string)
}

func SavePath(log *slog.Logger, saver Saver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http_server.handlers.save"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())))

		file, handler, err := r.FormFile("file")
		if err != nil {
			log.Error("Failed to get file from FormFile", sl.Error(err))
		}
		defer file.Close()

		pathFile := fmt.Sprintf("upload/%s", handler.Filename)

		dst, err := os.Create(pathFile)
		if err != nil {
			log.Error("Failed to create temp file in Dir", sl.Error(err))
		}
		defer dst.Close()
		
		if _, err := io.Copy(dst, file); err != nil {
			log.Error("Failed to copy file in Dir", sl.Error(err))
		}
		log.Info("File upload!")
	}

}
