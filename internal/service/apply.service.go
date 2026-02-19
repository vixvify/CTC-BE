package service

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repository"
	"time"

	"github.com/google/uuid"
)

type ApplyService struct {
	uow repository.UnitOfWork
}

func NewApplyService(uow repository.UnitOfWork) *ApplyService {
	return &ApplyService{uow: uow}
}

func (s *ApplyService) ApplyCamp(req dto.ApplyRequest, userID uuid.UUID) (dto.ApplyRequest, error) {
	err := s.uow.Do(func(team repository.TeamRepository, quiz repository.QuizRepository) error {

		t := models.Team{
			ID:        uuid.New(),
			Teamname:  req.Teamname,
			School:    req.School,
			UserID:    userID,
			CreatedAt: time.Now(),
		}

		if _, err := team.Create(t); err != nil {
			return err
		}

		q := models.Quiz{
			ID:        uuid.New(),
			Verified:  req.Verified,
			UserID:    userID,
			CreatedAt: time.Now(),
		}

		if _, err := quiz.Create(q); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return dto.ApplyRequest{}, err
	}

	return req, nil
}
