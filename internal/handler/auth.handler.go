package handler

import (
	"net/http"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/response"
	"server/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Signup(c *gin.Context) {
	var data dto.CreateUserRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	created, err := h.service.Signup(data)
	if err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.Created(c, dto.CreateUserResponse{
		UserDTO: mapper.ToUserResponse(created),
	})

}

func (h *AuthHandler) Login(c *gin.Context) {
	var data dto.LoginRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_, token, err := h.service.Login(data)
	if err != nil {
		response.Internal(c, err.Error())
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	response.OK(c, nil)
}

func (h *AuthHandler) Update(c *gin.Context) {
	var data dto.UpdateRequest

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

	updated, err := h.service.Update(userID, data)
	if err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.OK(c, dto.UpdateResponse{
		UserDTO: mapper.ToUserResponse(updated),
	})

}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var data dto.ChangePasswordRequest

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

	_, err = h.service.ChangePassword(userID, data)
	if err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.OK(c, nil)

}

func (h *AuthHandler) Delete(c *gin.Context) {
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

	err = h.service.Delete(userID)
	if err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.OK(c, nil)
}

func (h *AuthHandler) Me(c *gin.Context) {
	cookie, err := c.Request.Cookie("access_token")
	if err != nil {
		response.Unauthorized(c, "Access token not found")
		return
	}

	user, err := h.service.Me(cookie.Value)
	if err != nil {
		response.Unauthorized(c, "Unverified token")
		return
	}
	response.OK(c, dto.MeResponse{
		UserDTO: mapper.ToUserResponse(user),
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	response.OK(c, nil)
}
