package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	db *sql.DB
}

func NewSqliteStorage(filepath string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS links (
		code TEXT PRIMARY KEY,
		url TEXT NOT NULL
	);`

	if _, err = db.Exec(query); err != nil {
		return nil, err
	}

	return &SQLiteStorage{db: db}, nil
}

func (s *SQLiteStorage) Get(code string) (string, error) {
	var url string

	err := s.db.QueryRow("SELECT url FROM links WHERE code = ?", code).Scan(&url)
	if err == sql.ErrNoRows {
		return "", ErrLinkNotFound
	} else if err != nil {
		return "", err
	}

	return url, nil
}

func (s *SQLiteStorage) Put(code, url string) error {
	query := "INSERT INTO links (code, url) VALUES (?, ?)"
	_, err := s.db.Exec(query, code, url)
	return err
}