package elasticsearch

import (
	"errors"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
)

const serverUrl string = "http://localhost:9200"

func Ping() {
	_, _, err := CreateClient().Ping(serverUrl).Do(context.TODO())

	if err != nil {
		panic(err)
	}
}

func CreateClient() *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL(serverUrl))

	if err != nil {
		panic(err)
	}

	return client
}

func CreateIndex(mapping string, indexName string) error {
	client := CreateClient()

	exists, err := client.IndexExists(indexName).Do(context.TODO())

	if err != nil {
		return err
	}

	if !exists {
		response, err := client.CreateIndex(indexName).BodyString(mapping).Do(context.TODO())

		if err != nil {
			return err
		}

		if response == nil || !response.Acknowledged {
			return errors.New("Unable to create index")
		}
	}

	return nil
}

//func PutMapping(mapping string, indexName string) error {
// todo...
//	mappingResponse, err := client.PutMapping().Index(indexName).Type(indexName).BodyString(mapping).Do(context.TODO())
//
//	if err != nil {
//		return err
//	}
//
//	if mappingResponse == nil || !mappingResponse.Acknowledged {
//		return errors.New("Unable to put mapping")
//	}
//}

//func RemoveIndex(mapping string, indexName string) error {
// todo...
//}
