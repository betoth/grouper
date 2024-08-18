package output

import "grouper/application/domain"

type TopicPort interface {
	FindByID(topicID string) (*domain.Topic, error)
}
