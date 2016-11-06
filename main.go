package main

import (
	"github.com/apisearch/importer/routers"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server...")
	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
