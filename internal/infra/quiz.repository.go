package infra

import (
	"server/internal/models"
	"server/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuizRepoGorm struct {
	db *gorm.DB
}

func NewQuizRepoGorm(db *gorm.DB) repository.QuizRepository {
	return &QuizRepoGorm{db: db}
}

func (r *QuizRepoGorm) FindByID(id uuid.UUID) (models.Quiz, error) {
	var quiz models.Quiz
	err := r.db.Where("id = ?", id).First(&quiz).Error
	return quiz, err
}

func (r *QuizRepoGorm) Create(quiz models.Quiz) (models.Quiz, error) {
	err := r.db.Create(&quiz).Error
	return quiz, err
}

func (r *QuizRepoGorm) Update(id uuid.UUID, data models.Quiz) (models.Quiz, error) {
	if err := r.db.
		Model(&models.Quiz{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"verified": data.Verified,
			"video":    data.Video,
			"quiz_1":   data.Quiz_1,
			"quiz_2":   data.Quiz_2,
			"quiz_3":   data.Quiz_3,
			"quiz_4":   data.Quiz_4,
			"quiz_5":   data.Quiz_5,
		}).
		Error; err != nil {
		return models.Quiz{}, err
	}

	var quiz models.Quiz
	if err := r.db.First(&quiz, "id = ?", id).Error; err != nil {
		return models.Quiz{}, err
	}

	return quiz, nil
}
