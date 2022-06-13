package models

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	ID          int        `db:"id" json:"id"`
	Slug        string     `db:"slug" json:"slug"`
	CategoryID  *int       `db:"category_id" json:"categoryId"`
	Title       string     `db:"title" json:"title"`
	Body        string     `db:"body" json:"body"`
	CoverImage  string     `db:"cover_image" json:"coverImage"`
	Status      string     `db:"status" json:"status"`
	PublishedAt *time.Time `db:"published_at" json:"publishedAt"`
	CreatedBy   string     `db:"created_by" json:"createdBy"`
	CreatedAt   time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updatedAt"`
}

type ArticleDetail struct {
	ID          int        `db:"id" json:"id"`
	Slug        string     `db:"slug" json:"slug"`
	Category    *string    `db:"category" json:"category"`
	Title       string     `db:"title" json:"title"`
	Body        string     `db:"body" json:"body"`
	CoverImage  string     `db:"cover_image" json:"coverImage"`
	Status      string     `db:"status" json:"status"`
	Author      *string    `db:"author" json:"author"`
	PublishedAt *time.Time `db:"published_at" json:"publishedAt"`
	CreatedAt   time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updatedAt"`
}

type ESResponse struct {
	Message string `json:"message"`
}

type ArticleStore interface {
	CreateArticle(article *Article) (string, error)
	IndexingArticle(slug string) (*ESResponse, error)
}

type ArticleQueryStore interface {
	ArticleList(ginCtx *gin.Context) (*[]ArticleDetail, error)
}
