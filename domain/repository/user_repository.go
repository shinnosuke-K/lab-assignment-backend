package repository

import (
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
	"gorm.io/gorm"
)

type UerRepository interface {
	GetByID(db *gorm.DB, studentNum int) (model.User, error)
	UpdateGraduate(db *gorm.DB, graduate bool) error
	UpdateEntered(db *gorm.DB) error
}
