package handler

import (
	"net/http"
	"server/internal/dto"
	"server/internal/service"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"data":       nil,
			"status":     "error",
			"statusCode": 400})
		return
	}

	created, err := h.service.Signup(data)
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
		"statusCode": 201,
	})

}

func (h *AuthHandler) Login(c *gin.Context) {
	var data dto.LoginRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":       nil,
			"status":     "error",
			"statusCode": 400})
		return
	}

	user, token, err := h.service.Login(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":       nil,
			"status":     "error",
			"statusCode": 500})
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

	c.JSON(200, gin.H{
		"data": dto.LoginResponse{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		},
		"status":     "success",
		"statusCode": 200,
	})
}

func (h *AuthHandler) Me(c *gin.Context) {
	cookie, err := c.Request.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"data":       nil,
			"status":     "unauthorized",
			"statusCode": 401,
		})
		return
	}

	user, err := h.service.Me(cookie.Value)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"data":       nil,
			"status":     "unauthorized",
			"statusCode": 201,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       user,
		"status":     "success",
		"statusCode": 200,
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

	c.JSON(http.StatusOK, gin.H{
		"data":       nil,
		"statusCode": 200,
		"status":     "success",
	})
}
