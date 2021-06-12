package service

import (
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
	"github.com/shinnosuke-K/lab-assignment-backend/domain/repository"
	"github.com/shinnosuke-K/lab-assignment-backend/infrastructure/datastore"
)

type UserViewService interface {
	GetByID(input UserViewInput) (*UserViewOutput, error)
	UpdateGraduate(input UserViewInput)
	UpdateEntered(input UserViewInput)
}

type UserViewInput struct {
	Num      int
	Graduate bool
}

type UserViewOutput struct {
	User model.User
}

type UserViewServiceImpl struct {
	User repository.UerRepository
	Conn repository.DBConnector
}

func NewUserViewService(conn repository.DBConnector) *UserViewServiceImpl {
	return &UserViewServiceImpl{
		User: &datastore.UserStore{},
		Conn: conn,
	}
}

func (u *UserViewServiceImpl) GetByID(input UserViewInput) (*UserViewOutput, error) {

	panic("implement me")
}

func (u *UserViewServiceImpl) UpdateGraduate(input UserViewInput) {

	panic("implement me")
}

func (u *UserViewServiceImpl) UpdateEntered(input UserViewInput) {
	panic("implement me")
}
