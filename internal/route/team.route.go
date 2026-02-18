package route

import (
	"server/internal/handler"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterTeamRoutes(r *gin.RouterGroup, h *handler.TeamHandler) {
	r.Use(middleware.RateLimitMiddleware())
	r.GET("", h.GetTeams)
	r.GET("/:id", h.GetTeamByID)
	r.PUT("/:id", h.UpdateTeam)
}
