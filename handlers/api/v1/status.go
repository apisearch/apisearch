package v1

import (
	"github.com/apisearch/importer/handlers/response"
	"github.com/apisearch/importer/model/elasticsearch"
	"net/http"
)

func Healthz(w http.ResponseWriter, _ *http.Request) {
	elasticsearch.Ping()
	response.WriteOk(w, "Service is running")
}
