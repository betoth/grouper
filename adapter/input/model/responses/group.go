package responses

import "time"

type Group struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	UserName  string     `json:"user_name,omitempty"`
	Topic     GroupTopic `json:"topic_name"`
	CreatedAt time.Time  `json:"created_at"`
}

type GroupTopic struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	SubTopic GroupSubtopic `json:"subtopic"`
}

type GroupSubtopic struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
