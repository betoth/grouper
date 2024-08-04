package entity

import "time"

type GroupEntity struct {
	ID        string
	Name      string    `db:"name"`
	UserID    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}
