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
		return models.User{}, "", errors.New("invalid credentials")
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

func (s *AuthService) Me(token string) (dto.LoginResponse, error) {

	claims, err := util.VerifyAccessToken(token, s.jwtSecret)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	user, err := s.repo.FindByID(userID)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}
