package controller

import (
	"net/http"

	"github.com/shinnosuke-K/lab-assignment-backend/service"
)

type ProfViewController struct {
	Service service.ProfViewService
}

func NewProfViewController(s *service.ProfViewServiceImpl) *ProfViewController {
	return &ProfViewController{
		Service: s,
	}
}

func (p *ProfViewController) Handler(w http.ResponseWriter, r *http.Request) {

}
