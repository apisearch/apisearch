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
