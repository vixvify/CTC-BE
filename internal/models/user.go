package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Stats     string    `json:"stats"`
	CreatedAt time.Time `json:"created_at"`

	Team *Team `gorm:"foreignKey:UserID;references:ID"`
}
