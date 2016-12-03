package importer

import (
	"encoding/xml"
	"github.com/apisearch/importer/model/products"
	model "github.com/apisearch/importer/model/settings"
	"gopkg.in/olivere/elastic.v5"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func ImportXmlFiles() {
	var settings model.Settings
	var allSettings = []model.Settings{}
	var err error

	allSettings, err = settings.GetAll()

	if err != nil {
		log.Println("Unable to load stats!")
	} else {
		for _, s := range allSettings {
			log.Println("Importing user #" + s.UserId + " from url " + s.FeedUrl + "...")
			importXmlFile(s)
			log.Println("Import of user #" + s.UserId + " from url " + s.FeedUrl + " finished")
		}
	}
}

func importXmlFile(s model.Settings) error {
	var productList products.ProductList
	var err error

	resp, err := http.Get(s.FeedUrl)

	if err != nil {
		log.Println("Unable to get xml file!")

		return err
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("Unable to read xml file!")

		return err
	}

	defer resp.Body.Close()

	err = xml.Unmarshal(bytes, &productList)

	if err != nil {
		log.Println("Unable to unmarshall xml file!")

		return err
	}

	var firstUpdated string = time.Now().Format(products.DateFormat)
	var bulk *elastic.BulkProcessor

	bulk, err = products.BulkStart(s.UserId)

	for i, _ := range productList.ProductList {
		productList.ProductList[i].UserId = s.UserId
		productList.ProductList[i].Updated = time.Now().Format(products.DateFormat)
		productList.ProductList[i].BulkIndex(bulk)
	}

	products.DeleteOlderThan(firstUpdated)

	err = products.BulkFlush(bulk)

	return nil
}
