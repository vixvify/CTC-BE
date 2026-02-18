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
			"error": err.Error(),
		})
		return
	}

	userIDStr, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user id"})
		return
	}

	created, err := h.service.ApplyCamp(data, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, created)
}
