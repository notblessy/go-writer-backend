package handlers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	dao "github.com/notblessy/go-writter-backend/daos"
	"github.com/notblessy/go-writter-backend/models"
	"github.com/notblessy/go-writter-backend/utils"
	"github.com/segmentio/ksuid"
)

// Handler for create new article
func CreateArticle(ginCtx *gin.Context, publisher *amqp.Publisher, dao *dao.DAO) {
	var requestBody models.Article
	ginCtx.ShouldBindBodyWith(&requestBody, binding.JSON)

	response := &utils.Response{}

	generateSlug := ksuid.New()
	generateTitleSlug := strings.Split(strings.ToLower(requestBody.Title), " ")

	slug := fmt.Sprintf(`%s-%s`, strings.Join(generateTitleSlug, "-"), generateSlug)

	article := &models.Article{
		CategoryID: requestBody.CategoryID,
		Title:      requestBody.Title,
		Slug:       slug,
		Body:       requestBody.Body,
		CoverImage: requestBody.CoverImage,
		Status:     requestBody.Status,
		CreatedBy:  requestBody.CreatedBy,
	}

	body, _ := json.Marshal(article)

	err := publisher.Publish("create_article", message.NewMessage(watermill.NewUUID(), body))

	if err != nil {
		response.Message = err.Error()
		utils.ResponseError(ginCtx, response)
		return
	}

	response.Data = gin.H{"mq": "[x] Message Sent for create article"}
	utils.ResponseOK(ginCtx, response)
}

func FindArticleList(ginCtx *gin.Context, dao *dao.DAO) {
	properties, err := dao.ArticleQueryStore.ArticleList(ginCtx)

	response := new(utils.Response)

	if err != nil && strings.Contains(err.Error(), "error retrieving") {
		response.Message = err.Error()
		utils.ResponseNotFound(ginCtx, response)
		return
	}

	if err != nil && strings.Contains(err.Error(), "data not found") {
		response.Message = err.Error()
		utils.ResponseNotFound(ginCtx, response)
		return
	}

	response.Data = properties

	utils.ResponseOK(ginCtx, response)
}
