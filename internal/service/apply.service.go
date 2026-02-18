package service

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repository"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApplyService struct {
	db       *gorm.DB
	teamrepo repository.TeamRepository
	quizrepo repository.QuizRepository
}

func NewApplyService(db *gorm.DB, t repository.TeamRepository, q repository.QuizRepository) *ApplyService {
	return &ApplyService{db: db, teamrepo: t, quizrepo: q}
}

func (s *ApplyService) ApplyCamp(req dto.ApplyRequest, userID uuid.UUID) (dto.ApplyRequest, error) {
	err := s.db.Transaction(func(tx *gorm.DB) error {

		teamRepo := s.teamrepo.WithTx(tx)
		quizRepo := s.quizrepo.WithTx(tx)

		team := models.Team{
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

		if _, err := teamRepo.Create(team); err != nil {
			return err
		}

		quiz := models.Quiz{
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

		if _, err := quizRepo.Create(quiz); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return dto.ApplyRequest{}, err
	}

	return req, nil
}
