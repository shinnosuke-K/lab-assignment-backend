package repository

import (
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
)

type UerRepository interface {
	GetByID(db DBHandler, studentNum string) (*model.User, error)
	UpdateGraduate(db DBHandler, studentNum string, graduate bool) error
	UpdateEntered(db DBHandler, studentNum string) error
}
