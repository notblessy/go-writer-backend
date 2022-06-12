package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/notblessy/go-writter-backend/api/handlers"
	dao "github.com/notblessy/go-writter-backend/daos"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func Routes(router *gin.RouterGroup, ch *amqp.Channel, logger *zap.Logger, dao *dao.DAO) {
	router.POST("/categories", func(c *gin.Context) {
		handlers.CreateArticleCategory(c, ch, dao)
	})
}
