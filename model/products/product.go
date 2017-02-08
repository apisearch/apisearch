package products

import (
	"github.com/apisearch/apisearch/model/elasticsearch"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
)

type Product struct {
	Id          string `xml:"ITEM_ID" json:"id"`
	UserId      string `xml:"-" json:"userId"`
	Name        string `xml:"PRODUCTNAME" json:"name"`
	Description string `xml:"DESCRIPTION" json:"description"`
	Url         string `xml:"URL" json:"url"`
	Img         string `xml:"IMGURL" json:"img"`
	Price       int    `xml:"PRICE_VAT" json:"price"`
	Updated     string `xml:"-" json:"updated"`
}

type ProductList struct {
	ProductList []Product `xml:"SHOPITEM" json:"products"`
}

type SuggestedTerms struct {
	Terms []string `json:"terms"`
}

const (
	indexName     = "products"
	typeName      = "product"
	indexSettings = `{
		"settings": {
			"number_of_shards": 1,
			"number_of_replicas": 0,
			"analysis": {
				"filter": {
					"unique": {
						"type": "unique",
						"only_on_same_position": "false"
					},
					"pattern_replace": {
						"type": "pattern_replace",
						"preserve_original": 1,
						"pattern": "\\b(\\w{1,3})\\s+(\\w{1,3})\\b",
						"replacement": "$1$2"
					},
					"stopwords": {
						"type": "stop",
						"ignore_case": true,
						"stopwords": ["právě", "že", "_czech_"]
					},
					"hunspell": {
						"type": "hunspell",
						"locale": "cs_CZ",
						"dedup": true
					},
					"shingle": {
						"type": "shingle",
						"filter_token": "",
						"max_shingle_size": 3
					},
					"min_length": {
						"type": "length",
						"min": 2
					},
					"min_word_length": {
						"type": "length",
						"min": 4
					},
					"lowercase": {
						"type": "lowercase",
						"min": 4
					}
				},
				"analyzer": {
					"hunspell": {
						"filter": [
							"pattern_replace",
							"min_length",
							"hunspell",
							"icu_folding",
							"unique"
						],
						"tokenizer": "standard"
					},
					"icu": {
						"filter": [
							"icu_folding",
							"stopwords",
							"hunspell",
							"stopwords",
							"unique"
						],
						"tokenizer": "standard"
					},
					"shingle": {
						"filter": [
							"shingle",
							"pattern_replace",
							"min_length",
							"hunspell",
							"icu_folding",
							"unique"
						],
						"tokenizer": "standard"
					},
					"words": {
						"filter": [
							"min_word_length",
							"stopwords",
							"lowercase",
							"unique"
						],
						"tokenizer": "standard"
					}
				}
			}
		}
	}`
	typeSettings = `{
		"product": {
			"properties": {
				"id": {
					"type": "keyword"
				},
				"userId": {
					"type": "keyword"
				},
				"name": {
					"type": "keyword",
					"fields": {
						"hunspell": {
							"type": "text",
							"analyzer": "hunspell"
						},
						"icu": {
							"type": "text",
							"analyzer": "icu"
						},
						"shingle": {
							"type": "text",
							"analyzer": "shingle"
						},
						"words": {
							"type": "text",
							"analyzer": "words",
							"fielddata": true
						}
					}
				},
				"description": {
					"type": "keyword",
					"fields": {
						"hunspell": {
							"type": "text",
							"analyzer": "hunspell"
						},
						"icu": {
							"type": "text",
							"analyzer": "icu"
						},
						"shingle": {
							"type": "text",
							"analyzer": "shingle"
						}
					}
				},
				"url": {
					"type": "keyword"
				},
				"img": {
					"type": "keyword"
				},
				"price": {
					"type": "float"
				},
				"updated": {
					"type": "date",
					"format": "date_time_no_millis"
				}
			}
		}
	}`
)

func CreateIndex(force bool) error {
	var err error

	if force {
		err = elasticsearch.DeleteIndex(indexName)

		if err != nil {
			return err
		}
	}

	err = elasticsearch.CreateIndex(indexSettings, indexName)

	if err != nil {
		return err
	}

	return elasticsearch.PutMapping(typeSettings, indexName, typeName)
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

func BulkStart() (*elastic.BulkProcessor, error) {
	client := elasticsearch.CreateClient()

	return client.BulkProcessor().Name(indexName).Workers(4).Do()
}

func BulkFlush(bulk *elastic.BulkProcessor) error {
	bulk.Flush()

	return bulk.Close()
}

func DeleteOlderThan(updated string) error {
	client := elasticsearch.CreateClient()

	q := elastic.NewBoolQuery()
	q = q.MustNot(elastic.NewRangeQuery("updated").To(updated))

	_, err := client.DeleteByQuery().
		Index(indexName).
		Type(typeName).
		Query(q).
		Do(context.TODO())

	return err
}
