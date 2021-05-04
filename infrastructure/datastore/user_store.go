package datastore

import (
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
	"gorm.io/gorm"
)

type UserStore struct{}

func (*UserStore) GetByID(db *gorm.DB, studentNum int) (model.User, error) {
	return model.User{}, nil
}

func (*UserStore) UpdateGraduate(db *gorm.DB, graduate bool) error {
	return nil
}

func (*UserStore) UpdateEntered(db *gorm.DB) error {
	return nil
}
