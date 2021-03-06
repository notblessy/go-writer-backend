package dao

import (
	"github.com/jmoiron/sqlx"
	"github.com/notblessy/go-writter-backend/configs"
	"github.com/notblessy/go-writter-backend/models"
	"github.com/olivere/elastic/v7"
)

type DAO struct {
	*sqlx.DB
	*elastic.Client

	models.ArticleCategoryStore
	models.ArticleStore
	models.ArticleQueryStore
}

func NewStore() *DAO {
	db, err := configs.CreateConnection()
	if err != nil {
		panic(err)
	}

	esclient, err2 := configs.GetESClient()
	if err2 != nil {
		panic("Client fail ")
	}

	return &DAO{
		DB: db,

		ArticleCategoryStore: NewArticleCategoryStore(db),
		ArticleStore:         NewArticleStore(db, esclient),
		ArticleQueryStore:    NewArticleQueryStore(esclient),
	}
}
