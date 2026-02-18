package main

import (
	"os"
	"server/internal/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	PORT := os.Getenv("PORT")

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Cookie", "x-from"},
		AllowCredentials: true,
	}))

	r.Run(":" + PORT)

}
