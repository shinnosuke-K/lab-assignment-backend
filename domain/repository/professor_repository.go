package repository

import (
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
	"gorm.io/gorm"
)

type ProfRepository interface {
	GetAll(db *gorm.DB) ([]model.Professor, error)
	GetByID(db *gorm.DB, id string) (model.Professor, error)
}
