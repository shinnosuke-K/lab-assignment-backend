package controller

import (
	"net/http"

	"github.com/shinnosuke-K/lab-assignment-backend/pkg/service"
)

func respondError(w http.ResponseWriter, err interface{}) {
	switch err.(type) {
	case *service.UserError:

	}
}
