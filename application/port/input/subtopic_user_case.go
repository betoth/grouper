package input

import "grouper/application/domain"

type SubtopicService interface {
	FindByID(subtopicID string) (*domain.Subtopic, error)
}
