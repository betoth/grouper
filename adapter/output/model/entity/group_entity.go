package entity

import "time"

type GroupEntity struct {
	ID        string
	Name      string    `db:"name"`
	UserID    string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
}
