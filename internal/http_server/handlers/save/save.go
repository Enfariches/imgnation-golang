package save

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
	"github.com/skip2/go-qrcode"
)

// Наследования и модификация структуры Response
type Response struct {
	resp.Response
	Uuid string `json:"uuid,omitempty"` //Теги (метаданные) определяют представления в JSON
}

//go:generate mockery --name=Saver
type Saver interface {
	SaveIMG(path string) error
}

func SaveImage(log *slog.Logger, saver Saver) http.HandlerFunc {
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
			log.Error("Invalid body", sl.Error(err))
			render.JSON(w, r, resp.Error("Invalid body")) //Response invalid body
			return
		}
		extension, _ := strings.CutPrefix(infoFile, "image/")

		uuid, _ := uuid.NewV7()
		uuidStr := uuid.String()

		newFilePath := fmt.Sprintf("upload/%s.%s", uuidStr, extension)

		dst, err := os.Create(newFilePath)
		if err != nil {
			log.Error("Failed to create temp file in Dir", sl.Error(err))
			return
		}
		defer dst.Close()

		if _, err = io.Copy(dst, file); err != nil {
			log.Error("Failed to copy file in Dir", sl.Error(err))
			return
		}

		err = saver.SaveIMG(uuidStr)
		if err != nil {
			log.Error("Failed to Save to DB path", sl.Error(err))
			return
		}

		log.Info("File upload!")
		QRGenerate(log, uuidStr, extension)
		//render.Data(w, r, png) //Отправка QR-code в виде application/octet-stream
		render.JSON(w, r, Response{*resp.OK(), uuidStr})
	}
}

func QRGenerate(log *slog.Logger, uuid, extension string) {

	url := fmt.Sprintf("http://localhost:8080/api/img/%s?extension=%s", uuid, extension) //cfg.http_server.Address <- Вывод из КФГ
	err := qrcode.WriteFile(url, qrcode.Medium, 512, "qr2.png")
	if err != nil {
		log.Error("Failed to create qr", sl.Error(err))
	}
}
