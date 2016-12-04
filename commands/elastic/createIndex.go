package elastic

import (
	"github.com/apisearch/apisearch/model/products"
	"github.com/apisearch/apisearch/model/settings"
	"log"
)

func CreateIndex(force bool) {
	var err error

	err = settings.CreateIndex(force)

	if err != nil {
		log.Println("Unable to create index 'settings': " + err.Error())
	} else {
		log.Println("Create index 'settings': OK")
	}

	err = products.CreateIndex(force)

	if err != nil {
		log.Println("Unable to create index 'products': " + err.Error())
	} else {
		log.Println("Create index 'products': OK")
	}
}
