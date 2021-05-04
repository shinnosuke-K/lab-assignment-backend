package datastore

import (
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
	"gorm.io/gorm"
)

type AnswerStore struct{}

func (*AnswerStore) Store(db *gorm.DB, answers []model.Answer) error {
	return nil
}

func (*AnswerStore) Fix(db *gorm.DB, correctedAns []model.Answer) error {
	return nil
}

func (*AnswerStore) GetAnswersByID(db *gorm.DB, userID string) ([]model.Answer, error) {
	return nil, nil
}
