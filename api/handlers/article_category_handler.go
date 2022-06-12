package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/goccy/go-json"
	dao "github.com/notblessy/go-writter-backend/daos"
	"github.com/notblessy/go-writter-backend/models"
	"github.com/notblessy/go-writter-backend/utils"
	"github.com/segmentio/ksuid"
	"github.com/streadway/amqp"
)

// Handler for create new category
func CreateArticleCategory(ginCtx *gin.Context, ch *amqp.Channel, dao *dao.DAO) {
	var requestBody models.ArticleCategoryDTO
	ginCtx.ShouldBindBodyWith(&requestBody, binding.JSON)

	response := &utils.Response{}

	generateSlug := ksuid.New()

	slug := fmt.Sprintf(`%s-%s`, strings.ToLower(requestBody.Name), generateSlug)

	category := &models.ArticleCategoryDTO{
		MessageType: "create-category",
		Name:        requestBody.Name,
		Slug:        slug,
		Description: requestBody.Description,
	}

	q, err := ch.QueueDeclare(
		"ARTICLE-QUEUE",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		response.Message = err.Error()
		utils.ResponseError(ginCtx, response)
		return
	}

	body, _ := json.Marshal(category)

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	log.Printf(" [x] Sent %s", body)

	// categorySlug, err := dao.ArticleCategoryStore.CreateArticleCategory(category)

	if err != nil {
		response.Message = err.Error()
		utils.ResponseError(ginCtx, response)
		return
	}

	response.Data = gin.H{"mq": "[x] Message Sent"}
	utils.ResponseOK(ginCtx, response)
}

func String(body []byte) {
	panic("unimplemented")
}
