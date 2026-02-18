package main

import (
	"os"
	"server/internal/database"
	"server/internal/handler"
	"server/internal/infra"
	"server/internal/route"
	"server/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	jwtSecret := os.Getenv("JWT_SECRET")
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	authRepo := infra.NewAuthRepoGorm(database.DB)
	authService := service.NewAuthService(authRepo, jwtSecret)
	authHandler := handler.NewAuthHandler(authService)

	teamRepo := infra.NewTeamRepoGorm(database.DB)
	teamService := service.NewTeamService(teamRepo)
	teamHandler := handler.NewTeamHandler(teamService)

	quizRepo := infra.NewQuizRepoGorm(database.DB)
	quizService := service.NewQuizService(quizRepo)
	quizHandler := handler.NewQuizHandler(quizService)

	applyService := service.NewApplyService(database.DB, teamRepo, quizRepo)
	applyHandler := handler.NewApplyHandler(applyService)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Cookie", "x-from"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	route.RegisterAuthRoutes(api.Group("/auth"), authHandler)
	route.RegisterTeamRoutes(api.Group("/team"), teamHandler)
	route.RegisterQuizRoutes(api.Group("/quiz"), quizHandler)
	route.RegisterApplyRoutes(api.Group("/apply"), applyHandler)

	r.Run(":" + PORT)

}
