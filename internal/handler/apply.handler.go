package handler

import (
	"net/http"
	"server/internal/dto"
	"server/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ApplyHandler struct {
	service *service.ApplyService
}

func NewApplyHandler(s *service.ApplyService) *ApplyHandler {
	return &ApplyHandler{service: s}
}

func (h *ApplyHandler) ApplyCamp(c *gin.Context) {
	var data dto.ApplyRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":       nil,
			"status":     "error",
			"statusCode": 400,
		})
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

	created, err := h.service.ApplyCamp(data, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":       nil,
			"status":     "error",
			"statusCode": 500})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data":       created,
		"status":     "success",
		"statusCode": 201})
}
