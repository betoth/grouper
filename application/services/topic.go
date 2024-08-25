package services

import (
	"grouper/application/domain"
	"grouper/application/errors"
	"grouper/application/port/input"
	"grouper/application/port/output"
	"grouper/config/logger"

	"go.uber.org/zap"
)

func NewTopicService(topicRepository output.TopicPort) input.TopicService {
	return &topicService{
		topicRepository,
	}
}

type topicService struct {
	repository output.TopicPort
}

func (service *topicService) FindByID(topicID string) (*domain.Topic, error) {
	logger.Debug("Init FindByID service", zap.String("journey", "FindByID"))

	topicRepo, err := service.repository.FindByID(topicID)
	if err != nil {
		return nil, errors.HandleServiceError(err, "Group", "FindByID")
	}

	return topicRepo, nil
}
