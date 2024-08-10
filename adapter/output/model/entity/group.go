package entity

import "time"

type Group struct {
	ID         string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name       string    `gorm:"type:varchar(255);not null"`
	UserID     string    `gorm:"type:uuid;not null"`
	TopicID    string    `gorm:"type:uuid"`
	SubtopicID string    `gorm:"type:uuid"`
	CreatedAt  time.Time `gorm:"type:timestamptz;not null;default:now()"`
}
