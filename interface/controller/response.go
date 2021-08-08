package controller

import (
	"net/http"

	"github.com/shinnosuke-K/lab-assignment-backend/service"
)

func respondError(w http.ResponseWriter, err interface{}) {
	switch err.(type) {
	case *service.UserError:

	}
}
