package main

import (
	"github.com/apisearch/importer/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
