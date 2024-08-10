package entity

import "time"

type Subtopic struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	TopicID   string    `gorm:"type:uuid;not null"`
	Topic     Topic     `gorm:"foreignKey:TopicID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time `gorm:"type:timestamptz;not null;default:now()"`
}
