package postgres

import (
	"database/sql"
	"fmt"
	"img/internal/storage"

	pq "github.com/lib/pq"
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

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS images(id SERIAL PRIMARY KEY, 
   					path TEXT NOT NULL UNIQUE);`)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return &Storage{db: db}, nil

}

func (s *Storage) SaveIMG(path string) error {
	const op = "storage.postgres.SaveIMG"

	stmt, err := s.db.Prepare(`INSERT INTO images(path) VALUES ($1);`)
	if err != nil {
		return fmt.Errorf("%s error prepare, %w", op, err)
	}
	_, err = stmt.Exec(path)
	if err != nil {
		if errSql, ok := err.(*pq.Error); ok && errSql.Code == "23505" {
			return fmt.Errorf("%s, %w", op, storage.ErrPathExists)
		}
		return fmt.Errorf("%s, %w", op, err)
	}

	return nil
}
