package mapper

import (
	"server/internal/dto"
	"server/internal/models"
)

func ToUserDTO(u models.User) dto.UserDTO {
	return dto.UserDTO{
		Username: u.Username,
		Email:    u.Email,
		Stats:    u.Stats,
	}
}
