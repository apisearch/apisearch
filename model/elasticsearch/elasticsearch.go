package elasticsearch

import (
	"errors"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"os"
)

const (
	DateFormat             = "2006-01-02T15:04:05-07:00"
	serverUrl       string = "http://localhost:9200"
	serverUrlDocker string = "http://elasticsearch:9200"
)

func Ping() {
	var err error

	_, _, err = CreateClient().Ping(getEsUrl()).Do(context.TODO())

	if err != nil {
		panic(err)
	}
}

func CreateClient() *elastic.Client {
	var err error

	client, err := elastic.NewClient(elastic.SetURL(getEsUrl()), elastic.SetSniff(false))

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

func PutMapping(mapping string, indexName string, typeName string) error {
	client := CreateClient()
	response, err := client.PutMapping().Index(indexName).Type(typeName).BodyString(mapping).Do(context.TODO())

	if err != nil {
		return err
	}

	if response == nil || !response.Acknowledged {
		return errors.New("Unable to put mapping")
	}

	return nil
}

func DeleteIndex(indexName string) error {
	client := CreateClient()
	exists, err := client.IndexExists(indexName).Do(context.TODO())

	if err != nil {
		return err
	}

	if exists {
		_, err := client.DeleteIndex(indexName).Do(context.TODO())

		return err
	}

	return nil
}

func getEsUrl() string {
	if os.Getenv("DOCKER") == "true" {
		return serverUrlDocker
	} else {
		return serverUrl
	}
}
