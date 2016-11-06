package elasticsearch

import (
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
