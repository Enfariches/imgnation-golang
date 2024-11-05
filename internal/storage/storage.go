package storage

import "errors"

var (
	ErrPathExists = errors.New("path is has exists")
	ErrNotFoundUUID = errors.New("UUID is not found")
)
