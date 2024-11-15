package storage

import "errors"

var (
	ErrPathExists = errors.New("path is has exists")
	ErrNotFoundUUID = errors.New("UUID is not found")

	ErrRedisNotFoundOctet = errors.New("octet is not found in Redis")
)
