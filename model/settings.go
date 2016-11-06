package model

import (
	"encoding/json"
	"errors"
	"github.com/apisearch/importer/model/elasticsearch"
	"golang.org/x/net/context"
)

type Settings struct {
	UserId           string `json:"-"`
	FeedUrl          string `json:"feedUrl"`
	FeedFormat       string `json:"feedFormat"`
	DownloadInterval int    `json:"downloadInterval"`
}

const (
	indexName = "settings"
	typeName  = "settings"
	mapping   = `{
		"settings":{
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
	}`
)

func (s *Settings) Upsert() error {
	client := elasticsearch.CreateClient()

	exists, err := client.IndexExists(indexName).Do(context.TODO())

	if err != nil {
		return err
	}

	if !exists {
		response, err := client.CreateIndex(indexName).Do(context.TODO())

		if err != nil {
			return err
		}

		if response == nil || !response.Acknowledged {
			return errors.New("Unable to create index")
		}
	}

	mappingResponse, err := client.PutMapping().Index(indexName).Type(typeName).BodyString(mapping).Do(context.TODO())

	if err != nil {
		return err
	}

	if mappingResponse == nil || !mappingResponse.Acknowledged {
		return errors.New("Unable to put mapping")
	}

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
