package products

import (
	"github.com/apisearch/apisearch/model/elasticsearch"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"reflect"
)

func Search(_ int, query string) (*ProductList, error) {
	client := elasticsearch.CreateClient()

	matchQuery := elastic.NewMatchQuery("name.hunspell", query)

	res, err := client.Search(indexName).Query(matchQuery).Do(context.TODO())

	if err != nil {
		return nil, err
	}

	var ptype Product
	var productList *ProductList

	for _, item := range res.Each(reflect.TypeOf(ptype)) {
		productList.ProductList = append(productList.ProductList, item.(Product))
	}

	return productList, nil
}
