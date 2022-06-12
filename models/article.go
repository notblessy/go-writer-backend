package models

import "time"

type Article struct {
	ID          int        `db:"id" json:"id"`
	Slug        string     `db:"slug" json:"slug"`
	CategoryID  int        `db:"category_id" json:"categoryId"`
	Title       string     `db:"title" json:"title"`
	Body        string     `db:"body" json:"body"`
	CoverImage  string     `db:"cover_image" json:"coverImage"`
	Status      string     `db:"status" json:"status"`
	PublishedAt *time.Time `db:"published_at" json:"publishedAt"`
	CreatedAt   time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updatedAt"`
}

type ArticleStore interface {
	CreateArticle(article *Article) (string, error)
}
