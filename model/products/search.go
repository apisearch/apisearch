package products

import (
	"encoding/json"
	"errors"
	"github.com/apisearch/apisearch/model/elasticsearch"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"math"
)

func Search(userId string, query string, limit int) (ProductList, error) {
	var item Product
	var productList ProductList

	client := elasticsearch.CreateClient()

	boolQuery := elastic.NewBoolQuery()
	boolQuery.MinimumNumberShouldMatch(1)

	// fulltext
	matchQuery := elastic.NewMultiMatchQuery(query, "name.hunspell^3", "name.icu^3", "name.shingle", "name", "description")
	matchQuery.Type("most_fields")
	matchQuery.TieBreaker(0.3)

	// autocomplete
	matchPrefixQuery := elastic.NewPrefixQuery("name.icu", query)

	if len(query) > 3 {
		// typos
		fuzzyQuery := elastic.NewMatchQuery("name.hunspell", query)
		fuzzyQuery.Fuzziness("1")
		fuzzyQuery.Analyzer("hunspell")

		boolQuery.Should(matchQuery, matchPrefixQuery, fuzzyQuery)
	} else {
		boolQuery.Should(matchQuery, matchPrefixQuery)
	}

	userTermQuery := elastic.NewTermQuery("userId", userId)
	boolQuery.Must(userTermQuery)

	res, err := client.Search(indexName).Query(boolQuery).Size(limit).Do(context.TODO())

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
