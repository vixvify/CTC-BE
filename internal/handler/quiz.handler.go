package handler

import (
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/response"
	"server/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type QuizHandler struct {
	service *service.QuizService
}

func NewQuizHandler(s *service.QuizService) *QuizHandler {
	return &QuizHandler{service: s}
}

func (h *QuizHandler) GetQuizByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid quiz ID format")
		return
	}

	team, err := h.service.GetQuizByID(id)
	if err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.OK(c, mapper.ToQuizResponse(team))
}

func (h *QuizHandler) UpdateQuiz(c *gin.Context) {
	var updatedquiz dto.QuizRequest
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid quiz ID format")
		return
	}

	if err := c.ShouldBindJSON(&updatedquiz); err != nil {
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

	updated, err := h.service.UpdateQuiz(id, updatedquiz, userID)
	if err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.OK(c, mapper.ToQuizResponse(updated))
}
