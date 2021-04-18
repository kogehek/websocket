package store

import (
	"database/sql"
	"errors"
	"fmt"
	"websocket/crypto"
	"websocket/jwt"
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
	r.db.QueryRow(
		`INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id`,
		u.Email,
		hashPassword,
	).Scan(&u.ID)

	token := jwt.CrateToken(u.ID)
	const updateSQL = `UPDATE users SET token_jwt = $2 WHERE id = $1`
	if _, err := r.db.Exec(updateSQL, u.ID, token); err != nil {
		fmt.Println(err)
	}
	return nil
}

func (r *UserRepository) Auth(u *model.User) (*model.User, error) {
	var encrypted_password string
	userSql := `SELECT id, encrypted_password, token_jwt, email  FROM users WHERE email = $1`
	err := r.db.QueryRow(userSql, u.Email).Scan(&u.ID, &encrypted_password, &u.JWT, &u.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Doesn't correct email or password")
		}
		fmt.Println(err)
		return nil, err
	}
	if crypto.CheckPasswordHash(u.Password, encrypted_password) {
		u.Password = ""
		return u, nil
	}

	return nil, errors.New("Doesn't correct email or password")
}
