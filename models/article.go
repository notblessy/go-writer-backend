package models

type Article struct {
	ID          int     `db:"id" json:"id"`
	Slug        string  `db:"slug" json:"slug"`
	CategoryID  int     `db:"category_id" json:"categoryId"`
	Title       string  `db:"title" json:"title"`
	Body        string  `db:"body" json:"body"`
	CoverImage  string  `db:"cover_image" json:"coverImage"`
	Status      string  `db:"status" json:"status"`
	PublishedAt *string `db:"published_at" json:"publishedAt"`
	CreatedAt   string  `db:"created_at" json:"createdAt"`
	UpdatedAt   *string `db:"updated_at" json:"updatedAt"`
}
