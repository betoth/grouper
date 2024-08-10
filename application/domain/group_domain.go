package domain

import "time"

type GroupDomain struct {
	ID         string
	Name       string
	UserID     string
	TopicID    string
	SubtopicID string
	CreatedAt  time.Time
}
