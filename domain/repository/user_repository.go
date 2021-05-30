package repository

import (
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
)

type UerRepository interface {
	GetByID(db DBHandler, studentNum int) (*model.User, error)
	UpdateGraduate(db DBHandler, graduate bool) error
	UpdateEntered(db DBHandler) error
}
