package handlers

import (
	"fmt"
	resp "img/internal/lib/api/response"
	"img/internal/lib/logger/sl"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid"
)

// Наследования и модификация структуры Response
type Response struct {
	resp.Response
	Uuid string `json:"uuid,omitempty"` //Теги (метаданные) определяют представления в JSON
}

//go:generate mockery --name=Saver
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
			render.JSON(w, r, resp.Error("Invalid body")) //Response invalid body
			return
		}
		res, _ := strings.CutPrefix(infoFile, "image/")

		uuid, _ := uuid.NewV7()
		uuidStr := uuid.String()

		newFilePath := fmt.Sprintf("upload/%s.%s", uuidStr, res)

		dst, err := os.Create(newFilePath)
		if err != nil {
			log.Error("Failed to create temp file in Dir", sl.Error(err))
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			log.Error("Failed to copy file in Dir", sl.Error(err))
			return
		}

		err = saver.SavePath(uuidStr)
		if err != nil {
			log.Error("Failed to Save to DB path", sl.Error(err))
			return
		}

		log.Info("File upload!")
		render.JSON(w, r, Response{*resp.OK(), uuidStr})
	}
}
