package dto

type UserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Stats    string `json:"stats"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateUserResponse struct {
	UserDTO
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateRequest struct {
	UserDTO
}

type UpdateResponse struct {
	UserDTO
}

type ChangePasswordRequest struct {
	Password string `json:"password"`
}

type MeResponse struct {
	UserDTO
}
