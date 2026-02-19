package route

import (
	"server/internal/handler"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterTeamRoutes(r *gin.RouterGroup, h *handler.TeamHandler, jwtSecret string) {
	r.Use(middleware.RateLimitMiddleware())
	r.GET("", middleware.JWTAuth(jwtSecret), h.GetTeams)
	r.GET("/:id", middleware.JWTAuth(jwtSecret), h.GetTeamByID)
	r.PUT("/:id", middleware.JWTAuth(jwtSecret), h.UpdateTeam)
}
