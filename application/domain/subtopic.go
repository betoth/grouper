package domain

import "time"

type Subtopic struct {
	ID        string
	Name      string
	TopicID   string
	CreatedAt time.Time
}
