package repository

import (
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
	"gorm.io/gorm"
)

type AnswerRepository interface {
	Store(db *gorm.DB, answers []model.Answer) error
	Fix(db *gorm.DB, correctedAns []model.Answer) error
	GetAnswersByID(db *gorm.DB, userID string) ([]model.Answer, error)
}
