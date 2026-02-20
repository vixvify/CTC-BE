package dto

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Stats    string    `json:"stats"`
}

type UpdateRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Stats    string `json:"stats"`
}

type ChangePasswordRequest struct {
	Password string `json:"password"`
}
