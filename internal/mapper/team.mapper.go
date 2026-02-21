package mapper

import (
	"server/internal/dto"
	"server/internal/models"
)

func ToTeamResponse(t models.Team) dto.TeamResponse {
	return dto.TeamResponse{
		ID:       t.ID,
		Teamname: t.Teamname,
		School:   t.School,
		Call_1:   t.Call_1,
		Call_2:   t.Call_2,
		Relation: t.Relation,
		Name_1:   t.Name_1,
		Name_2:   t.Name_2,
		Name_3:   t.Name_3,
		Name_4:   t.Name_4,
	}
}

func ToTeamResponseList(teams []models.Team) []dto.TeamResponse {
	out := make([]dto.TeamResponse, 0, len(teams))
	for _, t := range teams {
		out = append(out, ToTeamResponse(t))
	}
	return out
}
