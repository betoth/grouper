package dto

import (
	"time"
)

type GroupDTO struct {
	ID        string
	Name      string
	UserName  string
	Topic     GroupTopic
	CreatedAt time.Time
}

type GroupTopic struct {
	ID       string
	Name     string
	Subtopic GroupSubtopic
}

type GroupSubtopic struct {
	ID   string
	Name string
}
