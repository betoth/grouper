package entity

import "time"

// TODO: Alterar tabelas no banco para ficar com o nome padrão
type Topic struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:timestamptz;not null;default:now()"`
}

func (Topic) TableName() string {
	return "topic"
}
