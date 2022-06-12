package main

import (
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	dao "github.com/notblessy/go-writter-backend/daos"
	"github.com/notblessy/go-writter-backend/models"
)

type createCategoryHandler struct {
	Dao *dao.DAO
}

func (c createCategoryHandler) Handler(msg *message.Message) ([]*message.Message, error) {
	var category models.ArticleCategory

	err := json.Unmarshal(msg.Payload, &category)

	if err != nil {
		return nil, err
	}

	categorySlug, err := c.Dao.ArticleCategoryStore.CreateArticleCategory(&category)

	if err != nil {
		return nil, err
	}

	log.Println("Create category handler received message:", msg.UUID, categorySlug)

	msg = message.NewMessage(watermill.NewUUID(), []byte("message produced by structHandler"))
	return message.Messages{msg}, nil
}
