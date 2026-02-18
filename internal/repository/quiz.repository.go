package repository

import (
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuizRepository interface {
	WithTx(tx *gorm.DB) QuizRepository
	FindByID(id uuid.UUID) (models.Quiz, error)
	Create(data models.Quiz) (models.Quiz, error)
	Update(id uuid.UUID, data models.Quiz) (models.Quiz, error)
}
