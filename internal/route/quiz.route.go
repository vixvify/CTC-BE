package route

import (
	"server/internal/handler"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterQuizRoutes(r *gin.RouterGroup, h *handler.QuizHandler) {
	r.Use(middleware.RateLimitMiddleware())
	r.GET("/:id", h.GetQuizByID)
	r.PUT("/:id", h.UpdateQuiz)
}
