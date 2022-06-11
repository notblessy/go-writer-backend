package models

import "time"

type ArticleCategory struct {
	ID          int        `db:"id" json:"id"`
	Slug        string     `db:"slug" json:"slug"`
	Name        string     `db:"name" json:"name"`
	Description string     `db:"description" json:"description"`
	IsActive    bool       `db:"is_active" json:"isActive"`
	PublishedAt *time.Time `db:"published_at" json:"publishedAt"`
	CreatedAt   time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updatedAt"`
}

type ArticleCategoryStore interface {
	CreateArticleCategory(category *ArticleCategory) (string, error)
	// DeleteArticleCategory(slug string) error
}
