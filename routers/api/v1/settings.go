package v1

import (
	"encoding/json"
	"github.com/apisearch/importer/errors"
	"github.com/apisearch/importer/model"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetSettingsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var userId int
	var err error

	if userId, err = strconv.Atoi(vars["userId"]); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// TODO: GET settings
	settings := model.Settings{UserId: userId, FeedUrl: "www.adboos.com/xml", FeedFormat: "heureka", Frequency: 3600}

	if settings.UserId > 0 {
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(settings); err != nil {
			panic(err)
		}

		return
	}

	w.WriteHeader(http.StatusNotFound)

	if err := json.NewEncoder(w).Encode(errors.Json{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func CreateSettings(w http.ResponseWriter, r *http.Request) {
	var settings model.Settings
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.Unmarshal(body, &settings); err != nil {
		w.WriteHeader(422)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// TODO: PUT settings

	w.WriteHeader(http.StatusNoContent)

	if err := json.NewEncoder(w).Encode(settings); err != nil {
		panic(err)
	}
}

func DeleteSettings(w http.ResponseWriter, r *http.Request) {
	var settings model.Settings
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.Unmarshal(body, &settings); err != nil {
		w.WriteHeader(422)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// TODO: DELETE settings

	w.WriteHeader(http.StatusNoContent)

	if err := json.NewEncoder(w).Encode(settings); err != nil {
		panic(err)
	}
}
