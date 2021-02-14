package store

import (
	"database/sql"
	"websocket/model"
)

type UserRepository struct {
	db *sql.DB
}

func userRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	err := validModel(u, r.db)
	if err != nil {
		return err
	}

	return r.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.Password,
	).Scan(&u.ID)
}
