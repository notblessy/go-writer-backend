package main

import (
	"os"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/notblessy/go-writter-backend/api/routes"
	dao "github.com/notblessy/go-writter-backend/daos"
	"github.com/notblessy/go-writter-backend/utils"

	ginzap "github.com/gin-contrib/zap"
)

var (
	amqpURI    = "amqp://guest:guest@localhost:5672/"
	amqplogger = watermill.NewStdLogger(false, false)
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

	amqpConfig := amqp.NewDurableQueueConfig(amqpURI)

	publisher, err := amqp.NewPublisher(amqpConfig, amqplogger)
	if err != nil {
		panic(err)
	}

	server.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	server.Use(ginzap.RecoveryWithZap(logger, true))

	routeGroup := server.Group("")
	routes.Routes(routeGroup, publisher, logger, dao)

	server.Run(":" + os.Getenv("PORT"))
}
