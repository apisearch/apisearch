package products

import (
	"github.com/apisearch/importer/model/elasticsearch"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
)

type Product struct {
	Id          string `xml:"ITEM_ID",json:"id"`
	UserId      string `xml:"-",json:"userId"`
	Name        string `xml:"PRODUCTNAME",json:"name"`
	Description string `xml:"DESCRIPTION",json:"description"`
	Url         string `xml:"URL",json:"url"`
	Img         string `xml:"IMGURL",json:"img"`
	Price       int    `xml:"PRICE_VAT",json:"price"`
	Updated     string `xml:"-",json:"updated"`
}

type ProductList struct {
	ProductList []Product `xml:"SHOPITEM"`
}

const (
	indexName = "products"
	typeName  = "product"
	mapping   = `{
		"settings":{
			"number_of_shards": 1,
			"number_of_replicas": 0
		},
		"mappings":{
			"product":{
				"properties":{
					"id":{
						"type": "string",
						"index": "not_analyzed"
					},
					"userId":{
						"type": "integer"
					},
					"name":{
						"type": "string",
						"index": "not_analyzed"
					},
					"description":{
						"type": "string",
						"index": "not_analyzed"
					},
					"url":{
						"type": "string",
						"index": "not_analyzed"
					},
					"img":{
						"type": "string",
						"index": "not_analyzed"
					},
					"price":{
						"type": "float"
					},
					"updated":{
						"type": "date",
						"format": "date_time_no_millis"
					}
				}
			}
		}
	}`
	DateFormat = "2006-01-02T15:04:05-07:00"
)

func CreateIndex() error {
	return elasticsearch.CreateIndex(mapping, indexName)
}

func (p *Product) Upsert() error {
	client := elasticsearch.CreateClient()

	response, err := client.
		Index().
		Index(indexName).
		Type(typeName).
		Id(p.UserId + "__" + p.Id).
		BodyJson(p).
		Do(context.TODO())

	if response == nil || err != nil {
		return err
	}

	return nil
}

func (p *Product) BulkIndex(bulk *elastic.BulkProcessor) {
	request := elastic.
		NewBulkIndexRequest().
		Index(indexName).
		Type(typeName).
		Id(p.UserId + "__" + p.Id).
		Doc(p)

	bulk.Add(request)
}

func BulkStart(userId string) (*elastic.BulkProcessor, error) {
	client := elasticsearch.CreateClient()

	return client.BulkProcessor().Name("index-products-" + userId).Workers(4).Do()
}

func BulkFlush(bulk *elastic.BulkProcessor) error {
	bulk.Flush()

	return bulk.Close()
}

func DeleteOlderThan(datetime string) error {
	// TODO: delete by query...

	return nil
}
