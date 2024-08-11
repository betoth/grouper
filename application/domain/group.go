package domain

import "time"

type Group struct {
	ID         string
	Name       string
	UserID     string
	TopicID    string
	SubtopicID string
	CreatedAt  time.Time
}
