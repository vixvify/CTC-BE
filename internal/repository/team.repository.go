package repository

import (
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamRepository interface {
	WithTx(tx *gorm.DB) TeamRepository
	FindAll() ([]models.Team, error)
	FindByID(id uuid.UUID) (models.Team, error)
	Create(data models.Team) (models.Team, error)
	Update(id uuid.UUID, data models.Team) (models.Team, error)
}
