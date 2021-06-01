package controller

import (
	"net/http"

	"github.com/shinnosuke-K/lab-assignment-backend/service"
)

type AnswerViewController struct {
	Service service.AnswerViewService
}

func NewAnsViewController(s *service.AnswerViewServiceImpl) *AnswerViewController {
	return &AnswerViewController{
		Service: s,
	}
}

func (a *AnswerViewController) Hander(w http.ResponseWriter, r *http.Request) {

}
