package repository

import (
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
)

type ProfRepository interface {
	GetAll(db DBHandler) ([]model.Professor, error)
	GetByID(db DBHandler, id string) (model.Professor, error)
}
