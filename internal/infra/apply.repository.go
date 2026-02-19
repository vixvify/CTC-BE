package infra

import (
	"server/internal/repository"

	"gorm.io/gorm"
)

type gormUOW struct {
	db *gorm.DB
}

func NewApplyRepository(db *gorm.DB) repository.ApplyRepository {
	return &gormUOW{db: db}
}

func (u *gormUOW) Do(fn func(repository.TeamRepository, repository.QuizRepository) error) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		teamRepo := NewTeamRepoGorm(tx)
		quizRepo := NewQuizRepoGorm(tx)

		return fn(teamRepo, quizRepo)
	})
}
