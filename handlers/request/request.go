package request

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func Read(r *http.Request, request interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &request); err != nil {
		return err
	}

	return nil
}

func GetVarFromRequest(r *http.Request, variable string) (string, error) {
	var value string
	var err error

	vars := mux.Vars(r)

	value = strings.TrimSpace(vars[variable])

	if value == "" {
		err = errors.New("Empty " + variable)
	}

	return value, err
}
