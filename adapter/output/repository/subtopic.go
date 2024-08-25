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

func NewSubtopicRepository(database *gorm.DB) output.SubtopicPort {
	return &subtopicRepository{
		database,
	}
}

type subtopicRepository struct {
	db *gorm.DB
}

func (repo *subtopicRepository) FindByID(subtopicID string) (*domain.Subtopic, error) {
	logger.Debug("Init FindByID in subtopicRepository", zap.String("journey", "FindSubtopicByID"))

	var subtopic entity.Subtopic

	result := repo.db.Where("id = ?", subtopicID).First(&subtopic)
	if result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {
			logger.Info("Subtopic notfound", zap.String("journey", "FindSubtopicByID"))
			return nil, customerror.NewBusinessError(customerror.BUSSINES_ERROR_SUBTOPIC_NOT_FOUND)
		}
		logger.Error("Error while trying search Subtopic", result.Error, zap.String("journey", "FindSubtopicByID"))
		return nil, result.Error
	}

	subtopicDomain := converter.ConverterSubtopicEntityToDomain(&subtopic)

	logger.Debug("Finish FindByID in subtopicRepository", zap.String("journey", "FindByID"))
	return &subtopicDomain, nil

}
