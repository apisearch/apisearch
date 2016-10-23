package v1

import (
	"net/http"
)

func Healthz(w http.ResponseWriter, _ *http.Request) {
	// TODO: check DB
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
