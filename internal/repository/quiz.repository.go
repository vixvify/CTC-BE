package repository

import (
	"server/internal/models"

	"github.com/google/uuid"
)

type QuizRepository interface {
	Create(data models.Quiz) (models.Quiz, error)
	Update(id uuid.UUID, data models.Quiz) (models.Quiz, error)
}
