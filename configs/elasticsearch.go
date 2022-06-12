package configs

import (
	"os"

	"github.com/olivere/elastic/v7"
)

func GetESClient() (*elastic.Client, error) {
	esHost := os.Getenv("ES_HOST_DEFAULT")

	client, err := elastic.NewClient(elastic.SetURL(esHost),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	return client, err

}
