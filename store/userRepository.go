package store

import (
	"database/sql"
	"errors"
	"fmt"
	"websocket/crypto"
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

	hashPassword, _ := crypto.HashPassword(u.Password)
	return r.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		hashPassword,
	).Scan(&u.ID)
}

func (r *UserRepository) Auth(u *model.User) (*model.User, error) {
	var encrypted_password string
	userSql := "SELECT id, encrypted_password FROM users WHERE email = $1"
	err := r.db.QueryRow(userSql, u.Email).Scan(&u.ID, &encrypted_password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Doesn't correct email or password")
		}
		fmt.Println(err)
		return nil, err
	}
	if crypto.CheckPasswordHash(u.Password, encrypted_password) {
		return u, nil
	}

	return nil, errors.New("Doesn't correct email or password")
}
