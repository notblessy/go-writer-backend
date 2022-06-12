package dao

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/notblessy/go-writter-backend/models"
)

type ArticleStore struct {
	*sqlx.DB
}

// Initialization article article store
func NewArticleStore(db *sqlx.DB) *ArticleStore {
	return &ArticleStore{DB: db}
}

// Insert data to table article article
func (store *ArticleStore) CreateArticle(article *models.Article) (string, error) {
	now := time.Now()
	article.CreatedAt = time.Now()

	columns := []string{
		"category_id",
		"title",
		"slug",
		"body",
		"cover_image",
		"created_at",
	}

	if article.Status == "PUBLISHED" {
		article.PublishedAt = &now
		columns = append(columns, "published_at")
	}

	query := fmt.Sprintf(`INSERT INTO articles (%s) VALUES (:%s)`, strings.Join(columns, ","), strings.Join(columns, ",:"))

	_, err := store.NamedExec(query, &article)

	return article.Slug, err
}
