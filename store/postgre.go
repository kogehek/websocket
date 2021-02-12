package store

import (
	"database/sql"
	"fmt"
)

// Store ...
type Store struct {
	db *sql.DB
}

// New create store
func New(dataBaseURL string) *Store {
	db, err := newDB(dataBaseURL)
	if err != nil {
		fmt.Println(err)
	}
	return &Store{
		db: db,
	}
}

func newDB(dataBaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataBaseURL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
