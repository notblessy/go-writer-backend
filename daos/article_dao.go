package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/notblessy/go-writter-backend/models"
	"github.com/olivere/elastic/v7"
)

type ArticleStore struct {
	*sqlx.DB
	*elastic.Client
}

// Initialization article article store
func NewArticleStore(db *sqlx.DB, es *elastic.Client) *ArticleStore {
	return &ArticleStore{DB: db, Client: es}
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
		"created_by",
	}

	if article.Status == "PUBLISHED" {
		article.PublishedAt = &now
		columns = append(columns, "published_at")
	}

	query := fmt.Sprintf(`INSERT INTO articles (%s) VALUES (:%s)`, strings.Join(columns, ","), strings.Join(columns, ",:"))

	_, err := store.NamedExec(query, &article)

	return article.Slug, err
}

func (store *ArticleStore) IndexingArticle(slug string) (*models.ESResponse, error) {
	var esResp models.ESResponse
	var ad models.ArticleDetail

	qn := `
	SELECT
		a.id, a.title, a.slug, a.body, a.cover_image, c.name "category",
		a.status, a.published_at, a.created_at, a.updated_at, a.created_by "author"
	FROM articles a
	LEFT JOIN article_categories c ON a.category_id = c.id
	WHERE a.slug = ?
	`

	err := store.DB.Get(&ad, qn, slug)

	if err != nil {
		return &esResp, err
	}

	CreateOrUpdateESArticle(store.Client, &ad, &esResp)

	return &esResp, err
}

func CreateOrUpdateESArticle(es *elastic.Client, pr *models.ArticleDetail, esResp *models.ESResponse) {
	var a models.ArticleDetail = *pr

	dataJSON, err := json.Marshal(a)
	js := string(dataJSON)
	_, err2 := es.Index().
		Index("articles").
		Id(a.Slug).
		BodyJson(js).
		Do(context.Background())

	if err2 != nil {
		panic(err)
	}

	success := "[Elastic][InsertArticle]Insertion Successful"
	esResp.Message = success
}
