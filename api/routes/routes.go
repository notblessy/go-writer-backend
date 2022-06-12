package routes

import (
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/gin-gonic/gin"
	"github.com/notblessy/go-writter-backend/api/handlers"
	dao "github.com/notblessy/go-writter-backend/daos"
	"go.uber.org/zap"
)

func Routes(router *gin.RouterGroup, publisher *amqp.Publisher, logger *zap.Logger, dao *dao.DAO) {
	router.POST("/categories", func(c *gin.Context) {
		handlers.CreateArticleCategory(c, publisher, dao)
	})

	router.GET("/articles", func(c *gin.Context) {
		handlers.FindArticleList(c, dao)
	})

	router.POST("/articles", func(c *gin.Context) {
		handlers.CreateArticle(c, publisher, dao)
	})
}
