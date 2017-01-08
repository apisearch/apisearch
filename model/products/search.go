package products

import (
	"encoding/json"
	"errors"
	"github.com/apisearch/apisearch/model/elasticsearch"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"math"
)

func Search(_ int, query string, limit int) (ProductList, error) {
	var item Product
	var productList ProductList

	client := elasticsearch.CreateClient()

	matchQuery := elastic.NewMatchQuery("name.hunspell", query)

	res, err := client.Search(indexName).Query(matchQuery).Size(limit).Do(context.TODO())

	if err != nil {
		return productList, err
	}

	productList.ProductList = make([]Product, int(math.Min(float64(limit), float64(res.TotalHits()))))

	i := 0

	for _, hit := range res.Hits.Hits {
		err = json.Unmarshal(*hit.Source, &item)

		if err != nil {
			return productList, errors.New("Unable to unmarshall product")
		}

		productList.ProductList[i] = item
		i++
	}

	return productList, nil
}
