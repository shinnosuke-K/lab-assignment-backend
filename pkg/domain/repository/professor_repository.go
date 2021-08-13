package repository

import (
	"github.com/shinnosuke-K/lab-assignment-backend/pkg/domain/model"
)

type ProfRepository interface {
	GetAll(DBHandler) ([]*model.Professor, error)
	GetByID(DBHandler, string) (*model.Professor, error)
}
