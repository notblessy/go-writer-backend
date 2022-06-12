package dao

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/notblessy/go-writter-backend/models"
)

type ArticleCategoryStore struct {
	*sqlx.DB
}

// Initialization article category store
func NewArticleCategoryStore(db *sqlx.DB) *ArticleCategoryStore {
	return &ArticleCategoryStore{DB: db}
}

// Insert data to table article category
func (store *ArticleCategoryStore) CreateArticleCategory(category *models.ArticleCategory) (string, error) {
	category.IsActive = true
	category.CreatedAt = time.Now()

	columns := []string{
		"name",
		"slug",
		"description",
		"is_active",
		"created_at",
	}
	query := fmt.Sprintf(`INSERT INTO article_categories (%s) VALUES (:%s)`, strings.Join(columns, ","), strings.Join(columns, ",:"))

	_, err := store.NamedExec(query, &category)

	return category.Slug, err
}
