package models

import "time"

type ArticleCategoryDTO struct {
	MessageType string     `json:"messageType"`
	ID          int        `db:"id" json:"id"`
	Slug        string     `db:"slug" json:"slug"`
	Name        string     `db:"name" json:"name"`
	Description string     `db:"description" json:"description"`
	IsActive    bool       `db:"is_active" json:"isActive"`
	CreatedAt   time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updatedAt"`
}

type ArticleCategoryStore interface {
	CreateArticleCategory(category *ArticleCategoryDTO) (string, error)
	// DeleteArticleCategory(slug string) error
}
