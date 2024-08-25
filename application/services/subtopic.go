package services

import (
	"grouper/application/domain"
	"grouper/application/port/input"
	"grouper/application/port/output"
	"grouper/config/logger"

	"go.uber.org/zap"
)

func NewSubtopicService(subtopicRepository output.SubtopicPort) input.SubtopicService {
	return &subtopicService{
		subtopicRepository,
	}
}

type subtopicService struct {
	repository output.SubtopicPort
}

func (service *subtopicService) FindByID(subtopicID string) (*domain.Subtopic, error) {
	logger.Debug("Init FindByID in subtopicService", zap.String("journey", "FindSubtopicByID"))

	subtopic, err := service.repository.FindByID(subtopicID)
	if err != nil {
		return nil, err
	}

	logger.Debug("Finish FindByID in subtopicService", zap.String("journey", "FindSubtopicByID"))
	return subtopic, nil
}
