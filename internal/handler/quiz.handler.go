package handler

import (
	"net/http"
	"server/internal/dto"
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
		c.JSON(400, gin.H{
			"data":       nil,
			"status":     "error",
			"statusCode": 400})
		return
	}

	team, err := h.service.GetQuizByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":       nil,
			"status":     "error",
			"statusCode": 500})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":       team,
		"status":     "success",
		"statusCode": 200,
	})
}

func (h *QuizHandler) UpdateQuiz(c *gin.Context) {
	var updatedquiz dto.ApplyRequest
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(400, gin.H{
			"data":       nil,
			"status":     "error",
			"statusCode": 400})
		return
	}

	if err := c.ShouldBindJSON(&updatedquiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":       nil,
			"status":     "error",
			"statusCode": 400})
		return
	}

	userIDStr, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"data":       nil,
			"status":     "error",
			"statusCode": 401})
		return
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":       nil,
			"status":     "error",
			"statusCode": 400})
		return
	}

	updated, err := h.service.UpdateQuiz(id, updatedquiz, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":       nil,
			"status":     "error",
			"statusCode": 500})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":       updated,
		"status":     "success",
		"statusCode": 200,
	})
}
