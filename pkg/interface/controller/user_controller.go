package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/shinnosuke-K/lab-assignment-backend/pkg/interface/presenter"
	"github.com/shinnosuke-K/lab-assignment-backend/pkg/service"
)

type UserViewController struct {
	Service service.UserViewService
}

func NewUserViewController(s *service.UserViewServiceImpl) *UserViewController {
	return &UserViewController{
		Service: s,
	}
}

type UserViewRequest struct {
	UserID string `json:"user_id"`
}

type UserViewResponse struct {
	UserID   string `json:"user_id"`
	Graduate bool   `json:"graduate"`
	Labs     []*Lab `json:"labs"`
}

type Lab struct {
	Name       string       `json:"name"`
	Professors []*Professor `json:"professors"`
}

type Professor struct {
	Name  string `json:"name"`
	Point int    `json:"point"`
}

func (u *UserViewController) GetUser(w http.ResponseWriter, r *http.Request) {

	uvr := UserViewRequest{}
	if err := decodeRequestBodyJSON(r, &uvr); err != nil {
		respondError(w, err)
		return
	}

	out, userErr := u.Service.GetByID(&service.UserViewInput{
		Num: uvr.UserID,
	})

	if userErr != nil {
		respondError(w, userErr)
		return
	}

	res := presenter.UserOKResponse(out.User, out.Professors, out.Answers)
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		respondError(w, errors.New("レスポンス作成中にエラーが発生しました"))
		return
	}
}
