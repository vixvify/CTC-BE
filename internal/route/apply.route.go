package route

import (
	"server/internal/handler"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterApplyRoutes(r *gin.RouterGroup, h *handler.ApplyHandler, jwtSecret string) {
	r.Use(middleware.RateLimitMiddleware())
	r.POST("/", middleware.JWTAuth(jwtSecret), h.ApplyCamp)
}
