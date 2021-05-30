package datastore

import (
	"github.com/pkg/errors"
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
	"github.com/shinnosuke-K/lab-assignment-backend/domain/repository"
)

type UserStore struct{}

func (*UserStore) GetByID(db repository.DBHandler, studentNum int) (*model.User, error) {

	q := `
			select * from users where student_num = ?
		`

	user := new(model.User)
	if err := db.QueryRowx(q, studentNum).StructScan(&user); err != nil {
		return &model.User{}, errors.Wrap(err, "UserStore.GetByID go error")
	}

	return user, nil
}

func (*UserStore) UpdateGraduate(db repository.DBHandler, graduate bool) error {
	return nil
}

func (*UserStore) UpdateEntered(db repository.DBHandler) error {
	return nil
}
