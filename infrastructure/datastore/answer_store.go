package datastore

import (
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
	"github.com/shinnosuke-K/lab-assignment-backend/domain/repository"
)

type AnswerStore struct{}

func (*AnswerStore) Store(db repository.DBHandler, answers []model.Answer) error {
	return nil
}

func (*AnswerStore) Fix(db repository.DBHandler, correctedAns []model.Answer) error {
	return nil
}

func (*AnswerStore) GetAnswersByID(db repository.DBHandler, userID string) ([]model.Answer, error) {
	return nil, nil
}
