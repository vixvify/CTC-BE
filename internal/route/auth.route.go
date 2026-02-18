package route

import (
	"server/internal/handler"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.RouterGroup, h *handler.AuthHandler) {
	r.Use(middleware.RateLimitMiddleware())
	r.POST("/signup", h.Signup)
	r.POST("/login", h.Login)
	r.GET("/me", h.Me)
	r.POST("/logout", h.Logout)
}
