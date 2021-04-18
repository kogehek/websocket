package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	DB             *sql.DB
	UserRepository *UserRepository
	RoomRepository *RoomRepository
}

// New create store
func New(dataBaseURL string) *Store {
	db, err := newDB(dataBaseURL)
	if err != nil {
		fmt.Println(err)
	}
	return &Store{
		DB:             db,
		UserRepository: userRepository(db),
		RoomRepository: roomRepository(db),
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
