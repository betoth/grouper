package domain

import "time"

type UserDomain struct {
	ID        string
	Name      string
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
}
