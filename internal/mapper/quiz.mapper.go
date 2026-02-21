package mapper

import (
	"server/internal/dto"
	"server/internal/models"
)

func ToQuizResponse(q models.Quiz) dto.QuizResponse {
	return dto.QuizResponse{
		ID:       q.ID,
		Verified: q.Verified,
		Video:    q.Video,
		Quiz_1:   q.Quiz_1,
		Quiz_2:   q.Quiz_2,
		Quiz_3:   q.Quiz_3,
		Quiz_4:   q.Quiz_4,
		Quiz_5:   q.Quiz_5,
	}
}
