package handlers

import (
	"fmt"
	"img/internal/lib/logger/sl"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/gofrs/uuid"
)

type Saver interface {
	SavePath(path string) error
}

func SavePath(log *slog.Logger, saver Saver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http_server.handlers.save"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())))

		file, header, err := r.FormFile("file")
		if err != nil {
			log.Error("Failed to get file from FormFile", sl.Error(err))
			return
		}
		defer file.Close()
		infoFile := header.Header.Get("Content-Type")

		if ok := strings.HasPrefix(infoFile, "image/"); !ok {
			log.Error("Invalid body") //Response invalid body
			return
		}
		res, _ := strings.CutPrefix(infoFile, "image/")

		uuid, _ := uuid.NewV7()
		uuidStr := uuid.String()

		filePath := fmt.Sprintf("upload/%s", header.Filename)
		newFilePath := fmt.Sprintf("upload/%s.%s", uuidStr, res)

		dst, err := os.Create(filePath)
		if err != nil {
			log.Error("Failed to create temp file in Dir", sl.Error(err))
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			log.Error("Failed to copy file in Dir", sl.Error(err))
			return
		}

		err = os.Rename(filePath, newFilePath)
		if err != nil {
			log.Error("Failed rename")
			return
		}

		err = saver.SavePath(uuidStr)
		if err != nil {
			log.Error("Failed to Save to DB path", sl.Error(err))
			return
		}
		log.Info("File upload!")
	}

}
