package routers

import (
	"github.com/apisearch/apisearch/log"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Queries     []string
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = log.Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler).
			Queries(route.Queries...)
	}

	return router
}
