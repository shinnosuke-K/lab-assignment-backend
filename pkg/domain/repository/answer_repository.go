package repository

import (
	"github.com/shinnosuke-K/lab-assignment-backend/pkg/domain/model"
)

type AnswerRepository interface {
	Store(DBHandler, []*model.Answer) error
	Fix(DBHandler, []*model.Answer) error
	GetAnswersByUserID(DBHandler, string) ([]*model.Answer, error)
}
