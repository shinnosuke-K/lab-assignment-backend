package repository

import (
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
)

type AnswerRepository interface {
	Store(db DBHandler, answers []model.Answer) error
	Fix(db DBHandler, correctedAns []model.Answer) error
	GetAnswersByID(db DBHandler, userID string) ([]model.Answer, error)
}
