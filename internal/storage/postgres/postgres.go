package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(StorageURL string) (*Storage, error) {
	const op = "storage.postgres.new"

	db, err := sql.Open("postgres", StorageURL)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS images(id INTEGER PRIMARY KEY, 
   					path TEXT);`)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return &Storage{db: db}, nil

}

func (s *Storage) SavePath(path string) {
	const op = "storage.postgres.SaveImg"
}
