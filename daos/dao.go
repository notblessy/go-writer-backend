package dao

import (
	"github.com/jmoiron/sqlx"
	"github.com/notblessy/go-writter-backend/configs"
	"github.com/notblessy/go-writter-backend/models"
)

type DAO struct {
	*sqlx.DB

	models.ArticleCategoryStore
}

func NewStore() *DAO {
	db, err := configs.CreateConnection()
	if err != nil {
		panic(err)
	}

	return &DAO{
		DB: db,

		ArticleCategoryStore: NewArticleCategoryStore(db),
	}
}
