package v1

import (
	"github.com/apisearch/apisearch/handlers/request"
	"github.com/apisearch/apisearch/handlers/response"
	model "github.com/apisearch/apisearch/model/settings"
	"net/http"
)

type userDetail struct {
	Id         string `json:"id"`
	Token      string `json:"token"`
	Email      string `json:"email"`
	FeedUrl    string `json:"feedUrl"`
	FeedFormat string `json:"feedFormat"`
}

func CreateSettings(w http.ResponseWriter, r *http.Request) {
	var settings model.Settings
	var err error
	var newUser model.NewUser

	if err = request.Read(r, &settings); err != nil {
		response.WriteError(w, "Failed to parse input", 400, err)

		return
	}

	if newUser, err = settings.Create(); err != nil {
		response.WriteError(w, "Unable to save settings", 400, err)

		return
	}

	response.WriteOkWithBody(w, newUser)
}

func UpdateSettings(w http.ResponseWriter, r *http.Request) {
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

	if err = settings.Update(userId); err != nil {
		response.WriteError(w, "Unable to save settings", 400, err)

		return
	}

	response.WriteOk(w, "Settings saved")
}

func GetSettingsByToken(w http.ResponseWriter, r *http.Request) {
	var s model.Settings
	var err error
	var token string
	var found bool

	if token, err = request.GetVarFromRequest(r, "token"); err != nil {
		response.WriteError(w, "User id not set", 400, err)

		return
	}

	if found, err = s.FindByToken(token); err != nil {
		response.WriteError(w, "Unable to get settings", 400, err)

		return
	}

	if !found {
		response.WriteError(w, "Settings not found", 400, err)

		return
	}

	response.WriteOkWithBody(w, userDetail{s.UserId, s.Token, s.Email, s.FeedUrl, s.FeedUrl})
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

	if found, err = settings.Remove(userId); err != nil {
		response.WriteError(w, "Unable to get settings", 400, err)

		return
	}

	if !found {
		response.WriteError(w, "Settings not found", http.StatusNotFound, err)

		return
	}

	response.WriteOk(w, "Settings deleted")
}
