package request

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
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
