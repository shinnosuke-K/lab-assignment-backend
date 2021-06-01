package datastore

import (
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
	"github.com/shinnosuke-K/lab-assignment-backend/domain/repository"
)

type ProfStore struct{}

func (*ProfStore) GetAll(db repository.DBHandler) ([]model.Professor, error) {
	return nil, nil
}

func (*ProfStore) GetByID(db repository.DBHandler, id string) (*model.Professor, error) {
	return &model.Professor{}, nil
}
