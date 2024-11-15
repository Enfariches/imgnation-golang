package save

import (
	resp "img/internal/lib/api/response"
	"img/internal/lib/logger/sl"
	"img/internal/lib/qr"
	"img/internal/lib/random"
	"img/internal/storage/s3"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
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

func SaveImage(addressEnv string, log *slog.Logger, db *s3.StorageS3) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http_server.handlers.save"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())))

		file, _, err := r.FormFile("file")
		if err != nil {
			log.Error("Failed to get file from FormFile", sl.Error(err))
			return
		}
		defer file.Close()

		key := random.RandStringByte(10)
		err = db.Save(log, file, key)

		if err != nil {
			log.Error("Failed to save file to S3", sl.Error(err))
			return
		}

		log.Info("File upload!")

		octet, err := qr.QRGenerate(addressEnv, log, key)
		if err != nil {
			log.Error("Failed to generate QR", sl.Error(err))
			return
		}
		render.Data(w, r, octet)
	}
}
