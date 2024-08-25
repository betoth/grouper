package output

import "grouper/application/domain"

type SubtopicPort interface {
	FindByID(subtopicID string) (*domain.Subtopic, error)
}
