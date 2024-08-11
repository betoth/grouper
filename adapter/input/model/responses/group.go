package responses

import "time"

type Group struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	UserName  string    `json:"user_name,omitempty"`
	Topic     Topic     `json:"topic_name"`
	CreatedAt time.Time `json:"created_at"`
}