package entity

import "time"

type UserGroup struct {
	ID       string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID   string    `gorm:"type:uuid;not null"`
	GroupID  string    `gorm:"type:uuid;not null"`
	JoinedAt time.Time `gorm:"type:timestamptz;not null;default:now()"`
}
