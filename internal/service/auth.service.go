package service

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repository"
	"server/internal/util"
	"time"

	"github.com/google/uuid"
)

type AuthService struct {
	repo      repository.AuthRepository
	jwtSecret string
}

func NewAuthService(r repository.AuthRepository, jwtSecret string) *AuthService {
	return &AuthService{
		repo:      r,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Signup(req dto.CreateUserRequest) (models.User, error) {
	hashed, err := util.HashPassword(req.Password)
	if err != nil {
		return models.User{}, err
	}
	user := models.User{
		ID:        uuid.New(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashed,
		Stats:     "pending",
		CreatedAt: time.Now(),
	}
	return s.repo.Signup(user)
}

func (s *AuthService) Login(req dto.LoginRequest) (models.User, string, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return models.User{}, "", err
	}
	if !util.CheckPassword(user.Password, req.Password) {
		return models.User{}, "", errors.New("wrong password")
	}
	token, err := util.GenerateAccessToken(
		user.ID.String(),
		s.jwtSecret,
	)
	if err != nil {
		return models.User{}, "", err
	}
	return user, token, nil
}

func (s *AuthService) Update(id uuid.UUID, req dto.UpdateRequest) (models.User, error) {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Stats:    req.Stats,
	}
	return s.repo.Update(id, user)
}

func (s *AuthService) ChangePassword(id uuid.UUID, req dto.ChangePasswordRequest) (models.User, error) {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Password: req.Password,
	}
	return s.repo.Update(id, user)
}

func (s *AuthService) Delete(id uuid.UUID) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *AuthService) Me(token string) (models.User, error) {

	claims, err := util.VerifyAccessToken(token, s.jwtSecret)
	if err != nil {
		return models.User{}, err
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return models.User{}, err
	}

	user, err := s.repo.FindByID(userID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
