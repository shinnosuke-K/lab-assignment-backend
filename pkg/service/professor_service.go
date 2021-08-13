package service

import (
	"github.com/shinnosuke-K/lab-assignment-backend/pkg/domain/model"
	"github.com/shinnosuke-K/lab-assignment-backend/pkg/domain/repository"
	"github.com/shinnosuke-K/lab-assignment-backend/pkg/infrastructure/datastore"
)

type ProfViewService interface {
	GetAll() (*ProfViewOutput, error)
	GetByID(input ProfViewInput) (*ProfViewOutput, error)
}

type ProfViewInput struct {
	ID string
}

type ProfViewOutput struct {
	Prof  *model.Professor
	Profs []model.Professor
}

type ProfViewServiceImpl struct {
	Prof repository.ProfRepository
	Conn repository.DBConnector
}

func NewProfViewService(conn repository.DBConnector) *ProfViewServiceImpl {
	return &ProfViewServiceImpl{
		Prof: &datastore.ProfStore{},
		Conn: conn,
	}
}

func (p *ProfViewServiceImpl) GetAll() (*ProfViewOutput, error) {
	panic("implement me")
}

func (p *ProfViewServiceImpl) GetByID(input ProfViewInput) (*ProfViewOutput, error) {
	panic("implement me")
}
