package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status      int    `json:"status"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, message string, status int, err error) {
	response := ErrorResponse{status, message, err.Error()}
	Write(w, status, response)
}

func WriteOk(w http.ResponseWriter, message string) {
	response := SuccessResponse{http.StatusOK, message}
	Write(w, http.StatusOK, response)
}

func WriteOkWithBody(w http.ResponseWriter, response interface{}) {
	Write(w, http.StatusOK, response)
}

func Write(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
