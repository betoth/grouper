package domain

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
}
