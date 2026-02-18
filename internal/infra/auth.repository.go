package infra

import (
	"server/internal/models"
	"server/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepoGorm struct {
	db *gorm.DB
}

func NewAuthRepoGorm(db *gorm.DB) repository.AuthRepository {
	return &AuthRepoGorm{db: db}
}

func (r *AuthRepoGorm) Signup(data models.User) (models.User, error) {
	err := r.db.Create(&data).Error
	return data, err
}

func (r *AuthRepoGorm) FindByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *AuthRepoGorm) FindByID(userID uuid.UUID) (models.User, error) {
	var user models.User

	err := r.db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *AuthRepoGorm) Update(id uuid.UUID, data models.User) (models.User, error) {
	if err := r.db.
		Model(&models.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"username": data.Username,
			"email":    data.Email,
			"stats":    data.Stats,
			"password": data.Password,
		}).
		Error; err != nil {
		return models.User{}, err
	}

	var blog models.User
	if err := r.db.First(&blog, "id = ?", id).Error; err != nil {
		return models.User{}, err
	}

	return blog, nil
}

func (r *AuthRepoGorm) Delete(userID uuid.UUID) error {
	err := r.db.Delete(&models.User{}, "id = ?", userID).Error
	if err != nil {
		return err
	}

	return nil
}
