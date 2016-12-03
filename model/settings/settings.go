package settings

import (
	"encoding/json"
	"github.com/apisearch/importer/model/elasticsearch"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"reflect"
)

type Settings struct {
	UserId           string `json:"userId"`
	FeedUrl          string `json:"feedUrl"`
	FeedFormat       string `json:"feedFormat"`
	DownloadInterval int    `json:"downloadInterval"`
}

const (
	indexName = "settings"
	typeName  = "setting"
	mapping   = `{
		"settings":{
			"number_of_shards": 1,
			"number_of_replicas": 0
		},
		"mappings":{
			"setting":{
				"properties":{
					"feedUrl":{
						"type":"string",
						"index": "no"
					},
					"feedFormat":{
						"type":"string",
						"index": "no"
					},
					"downloadInterval":{
						"type":"long"
					}
				}
			}
		}
	}`
)

func CreateIndex() error {
	return elasticsearch.CreateIndex(mapping, indexName)
}

func (s *Settings) Upsert() error {
	client := elasticsearch.CreateClient()

	response, err := client.Index().
		Index(indexName).
		Type(typeName).
		Id(s.UserId).
		BodyJson(s).
		Do(context.TODO())

	if response == nil || err != nil {
		return err
	}

	return nil
}

func (s *Settings) GetByUserId(userId string) (bool, error) {
	client := elasticsearch.CreateClient()

	res, err := client.Get().Index(indexName).Type(typeName).Id(userId).Do(context.TODO())

	if err != nil {
		return false, err
	}

	if res.Found != true || res.Source == nil {
		return false, nil
	}

	if err := json.Unmarshal(*res.Source, &s); err != nil {
		return false, err
	}

	return true, nil
}

func (s *Settings) RemoveByUserId(userId string) (bool, error) {
	client := elasticsearch.CreateClient()

	res, err := client.Delete().Index(indexName).Type(typeName).Id(userId).Do(context.TODO())

	if err != nil {
		return false, err
	}

	if res.Found != true {
		return false, nil
	}

	return true, nil
}

func (s *Settings) GetAll() ([]Settings, error) {
	client := elasticsearch.CreateClient()

	res, err := client.Search().Index(indexName).Query(elastic.NewMatchAllQuery()).Size(10000).Do(context.TODO())

	if err != nil {
		return nil, err
	}

	var ttyp Settings
	var result = []Settings{}

	for _, item := range res.Each(reflect.TypeOf(ttyp)) {
		if s, ok := item.(Settings); ok {
			result = append(result, s)
		}
	}

	return result, nil
}
