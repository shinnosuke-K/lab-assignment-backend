package service

import (
	"fmt"
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
	"github.com/shinnosuke-K/lab-assignment-backend/domain/repository"
	"github.com/shinnosuke-K/lab-assignment-backend/infrastructure/datastore"
	"net/http"
)

type UserViewService interface {
	GetByID(*UserViewInput) (*UserViewOutput, *UserError)
	UpdateGraduate(*UserViewInput) *UserError
	UpdateEntered(*UserViewInput) *UserError
}

type UserViewInput struct {
	Num      string
	Graduate bool
}

type UserViewOutput struct {
	User       *model.User
	Professors []*model.Professor
	Answers    []*model.Answer
}

type UserError struct {
	Status int
	Msg    string
}

type UserViewServiceImpl struct {
	User      repository.UerRepository
	Professor repository.ProfRepository
	Answer    repository.AnswerRepository
	Conn      repository.DBConnector
}

func NewUserViewService(conn repository.DBConnector) *UserViewServiceImpl {
	return &UserViewServiceImpl{
		User:      &datastore.UserStore{},
		Professor: &datastore.ProfStore{},
		Answer:    &datastore.AnswerStore{},
		Conn:      conn,
	}
}

func (u *UserViewServiceImpl) GetByID(input *UserViewInput) (*UserViewOutput, *UserError) {

	user, err := u.User.GetByID(u.Conn, input.Num)
	if err != nil {
		return nil, &UserError{
			Status: http.StatusInternalServerError,
			Msg:    fmt.Sprintf("id = %s のユーザは見つかりませでした。\nもう一度 id をご確認ください。\n", input.Num),
		}
	}

	professors, err := u.Professor.GetAll(u.Conn)
	if err != nil {
		return nil, &UserError{
			Status: http.StatusInternalServerError,
			Msg:    fmt.Sprintf("教授の情報を取得できませんでした。\n"),
		}
	}

	answers, err := u.Answer.GetAnswersByUserID(u.Conn, input.Num)
	if err != nil {
		return nil, &UserError{
			Status: http.StatusInternalServerError,
			Msg:    fmt.Sprintf("id = %s のユーザの回答を取得できませんでした。\nもう一度 id をご確認ください。\n", input.Num),
		}
	}

	return &UserViewOutput{
		User:       user,
		Professors: professors,
		Answers:    answers,
	}, nil
}

func (u *UserViewServiceImpl) UpdateGraduate(input *UserViewInput) *UserError {

	panic("implement me")
}

func (u *UserViewServiceImpl) UpdateEntered(input *UserViewInput) *UserError {
	panic("implement me")
}
