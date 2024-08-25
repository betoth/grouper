package repository

import (
	"grouper/adapter/output/converter"
	"grouper/adapter/output/model/entity"
	customerror "grouper/application/custom/custom-error"
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
	logger.Debug("Init FindByID in topicRepository", zap.String("journey", "FindTopicByID"))

	var topic entity.Topic

	result := repo.db.Where("id = ?", topicID).First(&topic)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			logger.Info("Topic notfound", zap.String("journey", "FindByID"))
			return nil, customerror.NewBusinessError(customerror.BUSSINES_ERROR_TOPIC_NOT_FOUND)
		}

		logger.Error("Error while trying search Topic", result.Error, zap.String("journey", "FindTopicByID"))
		return nil, result.Error
	}

	topicDomain := converter.ConverterTopicEntityToDomain(&topic)

	logger.Debug("Finish FindByID in topicRepository", zap.String("journey", "FindTopicByID"))
	return &topicDomain, nil

}
