package v1

import (
	"github.com/apisearch/apisearch/handlers/request"
	"github.com/apisearch/apisearch/handlers/response"
	"github.com/apisearch/apisearch/model/products"
	"net/http"
	"strconv"
)

const defaultLimit int = 20

func Search(w http.ResponseWriter, r *http.Request) {
	var err error
	var userIdParam string
	var userId int
	var query string
	var productList products.ProductList

	if userIdParam, err = request.GetVarFromRequest(r, "userId"); err != nil {
		response.WriteError(w, "User id not set", 400, err)

		return
	}

	if userId, err = strconv.Atoi(userIdParam); err != nil {
		response.WriteError(w, "Unable to parse user id", 400, err)

		return
	}

	if query, err = request.GetVarFromRequest(r, "query"); err != nil {
		response.WriteError(w, "Query string not set", 400, err)

		return
	}

	if productList, err = products.Search(userId, query, defaultLimit); err != nil {
		response.WriteError(w, "Search failed", 503, err)

		return
	}

	response.WriteOkWithBody(w, productList)
}
