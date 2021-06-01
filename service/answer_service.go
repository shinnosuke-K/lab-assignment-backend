package service

import (
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
	"github.com/shinnosuke-K/lab-assignment-backend/domain/repository"
	"github.com/shinnosuke-K/lab-assignment-backend/infrastructure/datastore"
)

type AnswerViewService interface {
	Store(input AnswerViewInput) error
	Fix(input AnswerViewInput) error
	GetAnswersByID(input AnswerViewInput) (*AnswerViewOutput, error)
}

type AnswerViewInput struct {
	Answers      []model.Answer
	CorrectedAns []model.Answer
	ID           string
}

type AnswerViewOutput struct {
	Ans []model.Answer
}

type AnswerViewServiceImpl struct {
	Answer repository.AnswerRepository
	Conn   repository.DBConnector
}

func NewAnsViewService(conn repository.DBConnector) *AnswerViewServiceImpl {
	return &AnswerViewServiceImpl{
		Answer: &datastore.AnswerStore{},
		Conn:   conn,
	}
}

func (a *AnswerViewServiceImpl) Store(input AnswerViewInput) error {
	panic("implement me")
}

func (a *AnswerViewServiceImpl) Fix(input AnswerViewInput) error {
	panic("implement me")
}

func (a *AnswerViewServiceImpl) GetAnswersByID(input AnswerViewInput) (*AnswerViewOutput, error) {
	panic("implement me")
}
