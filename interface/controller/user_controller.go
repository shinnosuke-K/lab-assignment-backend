package controller

import (
	"net/http"

	"github.com/shinnosuke-K/lab-assignment-backend/service"
)

type UserViewController struct {
	Service service.UserViewService
}

func NewUserViewController(s service.UserViewServiceImpl) *UserViewController {
	return &UserViewController{
		Service: &s,
	}
}

func (u *UserViewController) Handler(w http.ResponseWriter, r *http.Request) {
}
