package service

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repository"

	"github.com/google/uuid"
)

type TeamService struct {
	repo repository.TeamRepository
}

func NewTeamService(r repository.TeamRepository) *TeamService {
	return &TeamService{repo: r}
}

func (s *TeamService) GetTeams() ([]models.Team, error) {
	return s.repo.FindAll()
}

func (s *TeamService) GetTeamByID(id uuid.UUID) (models.Team, error) {
	return s.repo.FindByID(id)
}

func (s *TeamService) UpdateTeam(id uuid.UUID, req dto.TeamRequest, userID uuid.UUID) (models.Team, error) {
	teambyid, err := s.repo.FindByID(id)
	if err != nil {
		return models.Team{}, err
	}

	if userID != teambyid.UserID {
		return models.Team{}, errors.New("forbidden")
	}

	team := models.Team{
		Teamname: req.Teamname,
		School:   req.School,
		Call_1:   req.Call_1,
		Call_2:   req.Call_2,
		Relation: req.Relation,
		Name_1:   req.Name_1,
		Name_2:   req.Name_2,
		Name_3:   req.Name_3,
		Name_4:   req.Name_4,
	}
	return s.repo.Update(id, team)
}
