package entity

import "time"

// TODO: Alterar tabelas no banco para ficar com o nome padrão
type Subtopic struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	TopicID   string    `gorm:"type:uuid;not null"`
	Topic     Topic     `gorm:"foreignKey:TopicID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time `gorm:"type:timestamptz;not null;default:now()"`
}

func (Subtopic) TableName() string {
	return "subtopic"
}
