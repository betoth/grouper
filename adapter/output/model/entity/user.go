package entity

import "time"

type User struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"type:varchar(255)"`
	Email     string    `gorm:"type:varchar(255);unique"`
	Username  string    `gorm:"type:varchar(255);unique"`
	Password  string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
