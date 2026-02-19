package route

import (
	"server/internal/handler"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterQuizRoutes(r *gin.RouterGroup, h *handler.QuizHandler, jwtSecret string) {
	r.Use(middleware.RateLimitMiddleware())
	r.GET("/:id", middleware.JWTAuth(jwtSecret), h.GetQuizByID)
	r.PUT("/:id", middleware.JWTAuth(jwtSecret), h.UpdateQuiz)
}
