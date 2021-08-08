package repository

import (
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
)

type ProfRepository interface {
	GetAll(DBHandler) ([]*model.Professor, error)
	GetByID(DBHandler, string) (*model.Professor, error)
}
