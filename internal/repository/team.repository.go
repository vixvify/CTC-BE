package repository

import (
	"server/internal/models"

	"github.com/google/uuid"
)

type TeamRepository interface {
	FindAll() ([]models.Team, error)
	FindByID(id uuid.UUID) (models.Team, error)
	Create(data models.Team) (models.Team, error)
	Update(id uuid.UUID, data models.Team) (models.Team, error)
}
