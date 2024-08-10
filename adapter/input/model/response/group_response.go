package resp

import "time"

type GroupResponse struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	UserID     string    `json:"user_id,omitempty"`
	TopicID    string    `json:"topic_id"`
	SubTopicID string    `json:"subtopic_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type GroupResponse2 struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	UserName  string    `json:"user_name,omitempty"`
	Topic     Topic     `json:"topic_name"`
	CreatedAt time.Time `json:"created_at"`
}
