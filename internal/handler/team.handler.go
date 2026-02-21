package handler

import (
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/response"
	"server/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TeamHandler struct {
	service *service.TeamService
}

func NewTeamHandler(s *service.TeamService) *TeamHandler {
	return &TeamHandler{service: s}
}

func (h *TeamHandler) GetTeams(c *gin.Context) {
	teams, err := h.service.GetTeams()
	if err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.OK(c, mapper.ToTeamResponseList(teams))
}

func (h *TeamHandler) GetTeamByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid team ID format")
		return
	}

	team, err := h.service.GetTeamByID(id)
	if err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.OK(c, mapper.ToTeamResponse(team))
}

func (h *TeamHandler) UpdateTeam(c *gin.Context) {
	var updatedteam dto.TeamRequest
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid team ID format")
		return
	}

	if err := c.ShouldBindJSON(&updatedteam); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	userIDStr, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "User ID not found in context")
		return
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		response.BadRequest(c, "Invalid user ID format")
		return
	}

	updated, err := h.service.UpdateTeam(id, updatedteam, userID)
	if err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.OK(c, mapper.ToTeamResponse(updated))
}
