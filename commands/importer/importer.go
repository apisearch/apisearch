package importer

import (
	"encoding/xml"
	"github.com/apisearch/apisearch/model/elasticsearch"
	"github.com/apisearch/apisearch/model/products"
	model "github.com/apisearch/apisearch/model/settings"
	"gopkg.in/olivere/elastic.v5"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func ImportXmlFiles() {
	var settings model.Settings
	var allSettings = []model.Settings{}
	var err error
	var imported int
	var message string

	allSettings, err = settings.FindAll()

	if err != nil {
		log.Println("Unable to load stats!")
	} else {
		for _, s := range allSettings {
			log.Println("Importing user #" + s.UserId + " from url " + s.FeedUrl + "...")
			imported, err = importXmlFile(s)
			message = "Import of user #" + s.UserId + " from url " + s.FeedUrl

			if err != nil {
				log.Println(message + " failed: " + err.Error())
			} else {
				log.Println(message + " finished: stored " + strconv.Itoa(imported) + " products")
			}
		}
	}
}

func importXmlFile(s model.Settings) (int, error) {
	var productList products.ProductList
	var err error

	resp, err := http.Get(s.FeedUrl)

	if err != nil {
		log.Println("Unable to get xml file!")

		return 0, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("Unable to read xml file!")

		return 0, err
	}

	defer resp.Body.Close()

	err = xml.Unmarshal(bytes, &productList)

	if err != nil {
		log.Println("Unable to unmarshall xml file!")

		return 0, err
	}

	var firstUpdated string = time.Now().Format(elasticsearch.DateFormat)
	var bulk *elastic.BulkProcessor

	bulk, err = products.BulkStart()

	for i, _ := range productList.ProductList {
		productList.ProductList[i].UserId = s.UserId
		productList.ProductList[i].Updated = time.Now().Format(elasticsearch.DateFormat)
		productList.ProductList[i].BulkIndex(bulk)
	}

	products.DeleteOlderThan(firstUpdated)

	err = products.BulkFlush(bulk)

	return len(productList.ProductList), nil
}
