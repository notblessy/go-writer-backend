package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	ENV := os.Getenv("ENV")

	if ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.New()
	server.Run(":" + os.Getenv("PORT"))
}
