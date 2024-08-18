package input

import "grouper/application/domain"

type TopicService interface {
	FindByID(topicID string) (*domain.Topic, error)
}
