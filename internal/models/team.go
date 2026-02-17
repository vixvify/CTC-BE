package models

import (
	"time"

	"github.com/google/uuid"
)

type Team struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Teamname string    `json:"teamname"`
	School   string    `json:"school"`
	Call_1   string    `json:"call_1"`
	Call_2   string    `json:"call_2"`
	Relation string    `json:"relation"`
	Name_1   string    `json:"name_1"`
	Name_2   string    `json:"name_2"`
	Name_3   string    `json:"name_3"`
	Name_4   string    `json:"name_4"`

	UserID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex" json:"user_id"`
	User   *User     `gorm:"constraint:OnDelete:CASCADE;"`
	Quiz   *Quiz     `gorm:"foreignKey:TeamID;references:ID"`

	CreatedAt time.Time `json:"created_at"`
}
