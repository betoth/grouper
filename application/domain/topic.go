package domain

import "time"

type Topic struct {
	ID        string
	Name      string
	CreatedAt time.Time
}
