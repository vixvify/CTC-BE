package service

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repository"
	"time"

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

func (s *QuizService) CreateQuiz(req dto.ApplyRequest, userID uuid.UUID) (models.Quiz, error) {
	quiz := models.Quiz{
		ID:        uuid.New(),
		Verified:  req.Teamname,
		Video:     req.School,
		Quiz_1:    req.Call_1,
		Quiz_2:    req.Call_2,
		Quiz_3:    req.Relation,
		Quiz_4:    req.Name_1,
		Quiz_5:    req.Name_2,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	return s.repo.Create(quiz)
}

func (s *QuizService) UpdateQuiz(id uuid.UUID, req dto.ApplyRequest, userID uuid.UUID) (models.Quiz, error) {
	teambyid, err := s.repo.FindByID(id)
	if err != nil {
		return models.Quiz{}, err
	}

	if userID != teambyid.UserID {
		return models.Quiz{}, errors.New("forbidden")
	}

	team := models.Quiz{
		Verified: req.Teamname,
		Video:    req.School,
		Quiz_1:   req.Call_1,
		Quiz_2:   req.Call_2,
		Quiz_3:   req.Relation,
		Quiz_4:   req.Name_1,
		Quiz_5:   req.Name_2,
	}
	return s.repo.Update(id, team)
}
