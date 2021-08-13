package controller

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func decodeRequestBodyJSON(r *http.Request, u *UserViewRequest) error {
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return errors.New("")
	}
}
