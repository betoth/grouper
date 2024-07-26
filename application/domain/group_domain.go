package domain

import "time"

type GroupDomain struct {
	ID        string
	Name      string
	UserID    string
	CreatedAt time.Time
}
