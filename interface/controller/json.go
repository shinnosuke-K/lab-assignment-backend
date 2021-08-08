package controller

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

func decodeRequestBodyJSON(r *http.Request, u *UserViewRequest) error {
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return errors.New("")
	}
}
