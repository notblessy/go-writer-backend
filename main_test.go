package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/notblessy/go-writter-backend/api/handlers"
	dao "github.com/notblessy/go-writter-backend/daos"
	"github.com/notblessy/go-writter-backend/models"
)

func TestCreateArticleCategoryHandler(t *testing.T) {
	godotenv.Load()
	gin.SetMode(gin.TestMode)

	dao := dao.NewStore()

	r := gin.Default()

	amqpConfig := amqp.NewDurableQueueConfig(amqpURI)

	publisher, err := amqp.NewPublisher(amqpConfig, amqplogger)
	if err != nil {
		panic(err)
	}

	r.POST("/categories", func(c *gin.Context) {
		handlers.CreateArticleCategory(c, publisher, dao)
	})

	var category models.ArticleCategory
	category.Name = "Category Name Test"
	category.Description = "Description of category"
	category.CreatedBy = "Blessy"

	var reqBody bytes.Buffer
	err = json.NewEncoder(&reqBody).Encode(category)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "/articles", &reqBody)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestCreateArticleHandler(t *testing.T) {
	godotenv.Load()
	gin.SetMode(gin.TestMode)

	dao := dao.NewStore()

	r := gin.Default()

	amqpConfig := amqp.NewDurableQueueConfig(amqpURI)

	publisher, err := amqp.NewPublisher(amqpConfig, amqplogger)
	if err != nil {
		panic(err)
	}

	r.POST("/articles", func(c *gin.Context) {
		handlers.CreateArticle(c, publisher, dao)
	})

	var article models.Article
	article.Title = "Article Title Test"
	article.Body = "Body of article"
	article.CoverImage = "img-111.jpg"
	article.Status = "PUBLISHED"
	article.CreatedAt = time.Now()
	article.CreatedBy = "Blessy"

	var reqBody bytes.Buffer
	err = json.NewEncoder(&reqBody).Encode(article)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, "/articles", &reqBody)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestQueryArticleHandler(t *testing.T) {
	godotenv.Load()
	gin.SetMode(gin.TestMode)

	dao := dao.NewStore()

	r := gin.Default()
	r.GET("/articles", func(c *gin.Context) {
		handlers.FindArticleList(c, dao)
	})

	req, err := http.NewRequest(http.MethodGet, "/articles", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
