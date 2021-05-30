package service

import (
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
	"github.com/shinnosuke-K/lab-assignment-backend/domain/repository"
	"github.com/shinnosuke-K/lab-assignment-backend/infrastructure/datastore"
)

type UserViewService interface {
	GetByID(input UserViewInput) (UserViewOutput, Error)
	UpdateGraduate()
	UpdateEntered()
}

type UserViewInput struct {
	Num int
}

type UserViewOutput struct {
	User model.User
}

type Error struct {
	Message string
	Status  int
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

func (u *UserViewServiceImpl) GetByID(input UserViewInput) (UserViewOutput, Error) {
	panic("implement me")
}

func (u *UserViewServiceImpl) UpdateGraduate() {
	panic("implement me")
}

func (u *UserViewServiceImpl) UpdateEntered() {
	panic("implement me")
}
