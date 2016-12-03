package v1

import (
	"errors"
	"github.com/apisearch/importer/handlers/request"
	"github.com/apisearch/importer/handlers/response"
	model "github.com/apisearch/importer/model/settings"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func CreateSettings(w http.ResponseWriter, r *http.Request) {
	var settings model.Settings
	var err error
	var userId string

	if err = request.Read(r, &settings); err != nil {
		response.WriteError(w, "Failed to parse input", 422, err)

		return
	}

	if userId, err = getUserIdFromRequest(r); err != nil {
		response.WriteError(w, "User id not set", 422, err)

		return
	}

	settings.UserId = userId

	if err = settings.Upsert(); err != nil {
		response.WriteError(w, "Unable to save settings", 422, err)

		return
	}

	response.WriteOk(w, "Settings saved")
}

func GetSettingsById(w http.ResponseWriter, r *http.Request) {
	var settings model.Settings
	var err error
	var userId string
	var found bool

	if userId, err = getUserIdFromRequest(r); err != nil {
		response.WriteError(w, "User id not set", 422, err)

		return
	}

	if found, err = settings.GetByUserId(userId); err != nil {
		response.WriteError(w, "Unable to get settings", 422, err)

		return
	}

	if !found {
		response.WriteError(w, "Settings not found", http.StatusNotFound, err)

		return
	}

	response.WriteOkWithBody(w, settings)
}

func DeleteSettings(w http.ResponseWriter, r *http.Request) {
	var settings model.Settings
	var err error
	var userId string
	var found bool

	if userId, err = getUserIdFromRequest(r); err != nil {
		response.WriteError(w, "User id not set", 422, err)

		return
	}

	if found, err = settings.RemoveByUserId(userId); err != nil {
		response.WriteError(w, "Unable to get settings", 422, err)

		return
	}

	if !found {
		response.WriteError(w, "Settings not found", http.StatusNotFound, err)

		return
	}

	response.WriteOk(w, "Settings deleted")
}

func getUserIdFromRequest(r *http.Request) (string, error) {
	var userId string
	var err error

	vars := mux.Vars(r)

	userId = strings.TrimSpace(vars["userId"])

	if userId == "" {
		err = errors.New("Empty user id")
	}

	return userId, err
}
