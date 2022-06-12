package dao

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/notblessy/go-writter-backend/models"
	"github.com/olivere/elastic/v7"
)

type ArticleQueryStore struct {
	*elastic.Client
}

// Initialization article article store
func NewArticleQueryStore(es *elastic.Client) *ArticleQueryStore {
	return &ArticleQueryStore{Client: es}
}

func (store *ArticleQueryStore) ArticleList(ginCtx *gin.Context) (*[]models.ArticleDetail, error) {
	var adList []models.ArticleDetail

	s := elastic.NewBoolQuery()

	if ginCtx.Query("keyword") != "" {

		qn := fmt.Sprintf(`{"query_string": {"query": "(%s)","fields": [
			"title",
			"body"
		]}}`, ginCtx.Query("keyword"))
		qk := elastic.NewRawStringQuery(qn)
		s.Must(qk)
	}

	if ginCtx.Query("author") != "" {
		s.Filter(elastic.NewTermQuery("author.keyword", ginCtx.Query("author")))
	}

	searchService := store.Search().Index("articles").Query(s).Sort("createdAt", false)

	searchResult, err := searchService.Do(context.Background())

	if err != nil {
		return &[]models.ArticleDetail{}, fmt.Errorf("error retrieving data")
	}

	for _, hit := range searchResult.Hits.Hits {
		var arList *models.ArticleDetail

		err := json.Unmarshal(hit.Source, &arList)
		if err != nil {
			return &[]models.ArticleDetail{}, fmt.Errorf("error json unmarshal")
		}

		adList = append(adList, *arList)
	}

	return &adList, err
}
