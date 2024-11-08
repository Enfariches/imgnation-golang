package qr

import (
	"fmt"
	"img/internal/lib/logger/sl"
	"log/slog"

	"github.com/skip2/go-qrcode"
)

func QRGenerate(address string, log *slog.Logger, uuid, extension string) ([]byte, error) {

	url := fmt.Sprintf("http://localhost:8080/api/img/%s?extension=%s", uuid, extension) //cfg.http_server.Address <- Вывод из КФГ
	octet, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		log.Error("Failed to create qr", sl.Error(err))
	}
	return octet, err
}
