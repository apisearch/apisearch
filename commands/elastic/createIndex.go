package elastic

import (
	"github.com/apisearch/importer/model/products"
	"github.com/apisearch/importer/model/settings"
	"log"
)

func CreateIndex() {
	var err error

	err = settings.CreateIndex()

	if err != nil {
		log.Println("Unable to create index 'settings': " + err.Error())
	} else {
		log.Println("Create index 'settings': OK")
	}

	err = products.CreateIndex()

	if err != nil {
		log.Println("Unable to create index 'products': " + err.Error())
	} else {
		log.Println("Create index 'products': OK")
	}
}