package repository

import (
	"server/internal/models"

	"github.com/google/uuid"
)

type AuthRepository interface {
	Signup(data models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
	FindByID(userID uuid.UUID) (models.User, error)
	Update(id uuid.UUID, data models.User) (models.User, error)
	Delete(id uuid.UUID) error
}
