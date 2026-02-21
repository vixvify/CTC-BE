package dto

import "github.com/google/uuid"

type QuizResponse struct {
	ID       uuid.UUID `json:"id"`
	Verified string    `json:"verified"`
	Video    string    `json:"video"`
	Quiz_1   string    `json:"quiz_1"`
	Quiz_2   string    `json:"quiz_2"`
	Quiz_3   string    `json:"quiz_3"`
	Quiz_4   string    `json:"quiz_4"`
	Quiz_5   string    `json:"quiz_5"`
}

type QuizRequest struct {
	Verified string `json:"verified"`
	Video    string `json:"video"`
	Quiz_1   string `json:"quiz_1"`
	Quiz_2   string `json:"quiz_2"`
	Quiz_3   string `json:"quiz_3"`
	Quiz_4   string `json:"quiz_4"`
	Quiz_5   string `json:"quiz_5"`
}
