package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	dao "github.com/notblessy/go-writter-backend/daos"
	"github.com/notblessy/go-writter-backend/models"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	godotenv.Load()

	ENV := os.Getenv("ENV")

	if ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	dao := dao.NewStore()

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		"ARTICLE-QUEUE", // queue
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			_ = d
			var category models.ArticleCategoryDTO
			// err := json.Unmarshal(d.Body, &category)
			// if err != nil {
			// 	log.Printf(" [*] Error unmarshall data")
			// }
			category.Name = "Sport"
			category.Description = "Sport category test"
			category.Slug = "Sport-111"

			categorySlug, err := dao.ArticleCategoryStore.CreateArticleCategory(&category)
			if err != nil {
				log.Printf(" [*] Error create data: %v", err)
			}
			log.Printf(" [*] Data Created: %s", categorySlug)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
