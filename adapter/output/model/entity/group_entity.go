package entity

import "time"

type GroupEntity struct {
	ID         string
	Name       string    `db:"name"`
	UserID     string    `db:"user_id"`
	TopicID    string    `db:"topic_id"`
	SubTopicID string    `db:"subtopic_id"`
	CreatedAt  time.Time `db:"created_at"`
}
