package route

import (
	"server/internal/handler"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterApplyRoutes(r *gin.RouterGroup, h *handler.ApplyHandler) {
	r.Use(middleware.RateLimitMiddleware())
	r.POST("/", h.ApplyCamp)
}
