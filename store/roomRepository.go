package store

import (
	"database/sql"
	"websocket/model"
)

type RoomRepository struct {
	db *sql.DB
}

func roomRepository(db *sql.DB) *RoomRepository {
	return &RoomRepository{
		db: db,
	}
}

// Create ...
func (r *RoomRepository) Create(u *model.User, name string) error {
	var id int
	return r.db.QueryRow(
		"INSERT INTO rooms (user_id, name) VALUES ($1, $2) RETURNING id",
		u.ID,
		name,
	).Scan(&id)
}
