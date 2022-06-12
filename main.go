package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/notblessy/go-writter-backend/api/routes"
	dao "github.com/notblessy/go-writter-backend/daos"
	"github.com/notblessy/go-writter-backend/utils"
	"github.com/streadway/amqp"

	ginzap "github.com/gin-contrib/zap"
)

func main() {
	godotenv.Load()

	ENV := os.Getenv("ENV")

	if ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	logger, errlog := utils.NewLogger()
	if errlog != nil {
		panic(errlog)
	}

	defer logger.Sync()

	dao := dao.NewStore()
	server := gin.New()

	conn, errConn := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if errConn != nil {
		panic(errConn)
	}
	defer conn.Close()

	ch, errConn := conn.Channel()
	if errConn != nil {
		panic(errConn)
	}
	defer ch.Close()

	server.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	server.Use(ginzap.RecoveryWithZap(logger, true))

	routeGroup := server.Group("")
	routes.Routes(routeGroup, ch, logger, dao)

	server.Run(":" + os.Getenv("PORT"))
}
