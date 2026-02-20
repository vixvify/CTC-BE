package handler

import (
	"server/internal/dto"
	"server/internal/response"
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

	created, err := h.service.ApplyCamp(data, userID)
	if err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.Created(c, created)
}
