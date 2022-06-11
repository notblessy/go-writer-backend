package handlers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	dao "github.com/notblessy/go-writter-backend/daos"
	"github.com/notblessy/go-writter-backend/models"
	"github.com/notblessy/go-writter-backend/utils"
	"github.com/segmentio/ksuid"
)

// Handler for create new category
func CreateArticleCategory(ginCtx *gin.Context, dao *dao.DAO) {
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

	categorySlug, err := dao.ArticleCategoryStore.CreateArticleCategory(category)

	if err != nil {
		response.Message = err.Error()
		utils.ResponseError(ginCtx, response)
		return
	}

	response.Data = gin.H{"slug": categorySlug}
	utils.ResponseOK(ginCtx, response)
}
