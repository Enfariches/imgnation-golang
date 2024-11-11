package qr

import (
	"fmt"
	"img/internal/lib/logger/sl"
	"log/slog"

	"github.com/skip2/go-qrcode"
)

func QRGenerate(address string, log *slog.Logger, key string) ([]byte, error) {

	url := fmt.Sprintf("http://%s/api/img/%s", address, key) //cfg.http_server.Address <- Вывод из КФГ
	octet, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		log.Error("Failed to create qr", sl.Error(err))
		return nil, nil
	}
	return octet, err
}
