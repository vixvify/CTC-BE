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
			Call_1:    req.Call_1,
			Call_2:    req.Call_2,
			Relation:  req.Relation,
			Name_1:    req.Name_1,
			Name_2:    req.Name_2,
			Name_3:    req.Name_3,
			Name_4:    req.Name_4,
			UserID:    userID,
			CreatedAt: time.Now(),
		}

		if _, err := team.Create(t); err != nil {
			return err
		}

		q := models.Quiz{
			ID:        uuid.New(),
			Verified:  req.Verified,
			Video:     req.Video,
			Quiz_1:    req.Quiz_1,
			Quiz_2:    req.Quiz_2,
			Quiz_3:    req.Quiz_3,
			Quiz_4:    req.Quiz_4,
			Quiz_5:    req.Quiz_5,
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
