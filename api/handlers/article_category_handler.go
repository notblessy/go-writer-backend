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

// Handler for create new category
func CreateArticleCategory(ginCtx *gin.Context, publisher *amqp.Publisher, dao *dao.DAO) {
	var requestBody models.ArticleCategory
	ginCtx.ShouldBindBodyWith(&requestBody, binding.JSON)

	response := &utils.Response{}

	generateSlug := ksuid.New()

	slug := fmt.Sprintf(`%s-%s`, strings.ToLower(requestBody.Name), generateSlug)

	category := &models.ArticleCategory{
		Name:        requestBody.Name,
		Slug:        slug,
		Description: requestBody.Description,
	}

	body, _ := json.Marshal(category)

	err := publisher.Publish("create_category", message.NewMessage(watermill.NewUUID(), body))

	if err != nil {
		response.Message = err.Error()
		utils.ResponseError(ginCtx, response)
		return
	}

	response.Data = gin.H{"mq": "[x] Message Sent for create category"}
	utils.ResponseOK(ginCtx, response)
}
