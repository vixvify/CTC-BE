package models

import (
	"time"

	"github.com/google/uuid"
)


type Quiz struct {
	ID        	uuid.UUID 	`gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Verified    string		`json:"verified"`
	Video    	string		`json:"video"`
	Quiz_1    	string		`json:"quiz_1"`
	Quiz_2    	string		`json:"quiz_2"`
	Quiz_3    	string		`json:"quiz_3"`
	Quiz_4    	string		`json:"quiz_4"`
	Quiz_5    	string		`json:"quiz_5"`
	TeamID    	uuid.UUID	`json:"team_id"`
	CreatedAt 	time.Time 	`json:"created_at"`
}