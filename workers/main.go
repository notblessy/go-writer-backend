package main

import (
	"context"
	"os"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	dao "github.com/notblessy/go-writter-backend/daos"
)

var (
	amqpURI = "amqp://guest:guest@localhost:5672/"
	logger  = watermill.NewStdLogger(false, false)
)

func main() {
	godotenv.Load()

	ENV := os.Getenv("ENV")

	if ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	dao := dao.NewStore()

	amqpConfig := amqp.NewDurableQueueConfig(amqpURI)

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	// SignalsHandler will gracefully shutdown Router when SIGTERM is received.
	// You can also close the router by just calling `r.Close()`.
	router.AddPlugin(plugin.SignalsHandler)

	// Router level middleware are executed for every message sent to the router
	router.AddMiddleware(
		// CorrelationID will copy the correlation id from the incoming message's metadata to the produced messages
		middleware.CorrelationID,

		// The handler function is retried if it returns an error.
		// After MaxRetries, the message is Nacked and it's up to the PubSub to resend it.
		middleware.Retry{
			MaxRetries:      3,
			InitialInterval: time.Millisecond * 100,
			Logger:          logger,
		}.Middleware,

		// Recoverer handles panics from handlers.
		// In this case, it passes them as errors to the Retry middleware.
		middleware.Recoverer,
	)

	subscriber, err := amqp.NewSubscriber(amqpConfig, logger)
	if err != nil {
		panic(err)
	}

	publisher, err := amqp.NewPublisher(amqpConfig, logger)
	if err != nil {
		panic(err)
	}

	router.AddHandler(
		"handler_create_category",
		"create_category",
		subscriber,
		"create_category_processed",
		publisher,
		createCategoryHandler{Dao: dao}.Handler,
	)

	router.AddHandler("handler_create_article",
		"create_article",
		subscriber,
		"create_article_processed",
		publisher,
		createArticleHandler{Dao: dao}.Handler,
	)

	if err := router.Run(context.Background()); err != nil {
		panic(err)
	}
}
