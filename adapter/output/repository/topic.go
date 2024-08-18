package repository

import (
	"grouper/adapter/output/converter"
	"grouper/adapter/output/model/entity"
	"grouper/application/domain"
	"grouper/application/port/output"
	"grouper/config/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewTopicRepository(database *gorm.DB) output.TopicPort {
	return &topicRepository{
		database,
	}
}

type topicRepository struct {
	db *gorm.DB
}

func (repo *topicRepository) FindByID(topicID string) (*domain.Topic, error) {
	logger.Debug("Init FindByID repository", zap.String("journey", "FindByID"))

	var topic entity.Topic

	result := repo.db.Where("id = ?", topicID).First(&topic)
	if result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {
			logger.Info("Topic notfound", zap.String("journey", "FindByID"))
			return nil, result.Error
		}
		logger.Error("Error while trying search Topic", result.Error, zap.String("journey", "FindByID"))
		return nil, result.Error
	}

	topicDomain := converter.ConverterTopicEntityToDomain(&topic)

	logger.Debug("Finish FindByID repository", zap.String("journey", "FindByID"))
	return &topicDomain, nil

}
