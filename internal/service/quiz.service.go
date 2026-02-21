package service

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repository"

	"github.com/google/uuid"
)

type QuizService struct {
	repo repository.QuizRepository
}

func NewQuizService(r repository.QuizRepository) *QuizService {
	return &QuizService{repo: r}
}

func (s *QuizService) GetQuizByID(id uuid.UUID) (models.Quiz, error) {
	return s.repo.FindByID(id)
}

func (s *QuizService) UpdateQuiz(id uuid.UUID, req dto.QuizRequest, userID uuid.UUID) (models.Quiz, error) {
	teambyid, err := s.repo.FindByID(id)
	if err != nil {
		return models.Quiz{}, err
	}

	if userID != teambyid.UserID {
		return models.Quiz{}, errors.New("forbidden")
	}

	quiz := models.Quiz{
		Verified: req.Verified,
		Video:    req.Video,
		Quiz_1:   req.Quiz_1,
		Quiz_2:   req.Quiz_2,
		Quiz_3:   req.Quiz_3,
		Quiz_4:   req.Quiz_4,
		Quiz_5:   req.Quiz_5,
	}
	return s.repo.Update(id, quiz)
}
