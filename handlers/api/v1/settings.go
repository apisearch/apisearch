package v1

import (
	"github.com/apisearch/apisearch/handlers/request"
	"github.com/apisearch/apisearch/handlers/response"
	model "github.com/apisearch/apisearch/model/settings"
	"net/http"
)

func CreateSettings(w http.ResponseWriter, r *http.Request) {
	var settings model.Settings
	var err error
	var userId string

	if err = request.Read(r, &settings); err != nil {
		response.WriteError(w, "Failed to parse input", 400, err)

		return
	}

	if userId, err = request.GetVarFromRequest(r, "userId"); err != nil {
		response.WriteError(w, "User id not set", 400, err)

		return
	}

	settings.UserId = userId

	if err = settings.Upsert(); err != nil {
		response.WriteError(w, "Unable to save settings", 400, err)

		return
	}

	response.WriteOk(w, "Settings saved")
}

func GetSettingsById(w http.ResponseWriter, r *http.Request) {
	var settings model.Settings
	var err error
	var userId string
	var found bool

	if userId, err = request.GetVarFromRequest(r, "userId"); err != nil {
		response.WriteError(w, "User id not set", 400, err)

		return
	}

	if found, err = settings.GetByUserId(userId); err != nil {
		response.WriteError(w, "Unable to get settings", 400, err)

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

	if userId, err = request.GetVarFromRequest(r, "userId"); err != nil {
		response.WriteError(w, "User id not set", 400, err)

		return
	}

	if found, err = settings.RemoveByUserId(userId); err != nil {
		response.WriteError(w, "Unable to get settings", 400, err)

		return
	}

	if !found {
		response.WriteError(w, "Settings not found", http.StatusNotFound, err)

		return
	}

	response.WriteOk(w, "Settings deleted")
}
