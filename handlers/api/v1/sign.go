package v1

import (
	"github.com/apisearch/apisearch/handlers/request"
	"github.com/apisearch/apisearch/handlers/response"
	model "github.com/apisearch/apisearch/model/settings"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	var formData model.SignInData
	var err error
	var newUser model.NewUser

	if err = request.Read(r, &formData); err != nil {
		response.WriteError(w, "Failed to parse input", 400, err)

		return
	}

	if newUser, err = model.SignIn(formData); err != nil {
		response.WriteError(w, "Unable to sign in", 400, err)

		return
	}

	response.WriteOkWithBody(w, newUser)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	var err error
	var token string
	var found bool
	var settings model.Settings

	if token, err = request.GetVarFromRequest(r, "token"); err != nil {
		response.WriteError(w, "User id not set", 400, err)

		return
	}

	if found, err = settings.FindByToken(token); err != nil {
		response.WriteError(w, "Unable to get settings", 400, err)

		return
	}

	if !found {
		response.WriteError(w, "No user found", 400, err)

		return
	}

	if err = settings.SignOut(); err != nil {
		response.WriteError(w, "Unable to sign in", 400, err)

		return
	}

	response.WriteOk(w, "Signed out")
}
