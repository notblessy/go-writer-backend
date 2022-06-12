package main

import (
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	dao "github.com/notblessy/go-writter-backend/daos"
	"github.com/notblessy/go-writter-backend/models"
)

type createArticleHandler struct {
	Dao *dao.DAO
}

func (c createArticleHandler) Handler(msg *message.Message) ([]*message.Message, error) {
	var article models.Article

	err := json.Unmarshal(msg.Payload, &article)

	if err != nil {
		return nil, err
	}

	articleSlug, err := c.Dao.ArticleStore.CreateArticle(&article)

	if err != nil {
		return nil, err
	}

	es, err := c.Dao.ArticleStore.IndexingArticle(articleSlug)

	if err != nil {
		return nil, err
	}

	log.Println("Create article handler received message:", msg.UUID, articleSlug)
	log.Println("Success indexing to ES", msg.UUID, es)

	msg = message.NewMessage(watermill.NewUUID(), []byte("message produced by structHandler"))
	return message.Messages{msg}, nil
}
