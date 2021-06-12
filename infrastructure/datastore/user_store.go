package datastore

import (
	"github.com/pkg/errors"
	"github.com/shinnosuke-K/lab-assignment-backend/config/db"
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
	"github.com/shinnosuke-K/lab-assignment-backend/domain/repository"
)

type UserStore struct{}

func (*UserStore) GetByID(db repository.DBHandler, studentNum int) (*model.User, error) {

	q := `
			select * from users where student_num = ?
		`

	var user model.User
	if err := db.QueryRowx(q, studentNum).StructScan(&user); err != nil {
		return nil, errors.Wrap(err, "UserStore.GetByID go error")
	}

	return &user, nil
}

func (*UserStore) UpdateGraduate(db repository.DBHandler, studentNum int, graduate bool) error {

	if !exitsRecord(studentNum) {
		return errors.Errorf("no record with student_num = %d \n", studentNum)
	}

	q := `
		update users set graduate = ? where student_num = ?
		`

	if _, err := db.Exec(q, graduate, studentNum); err != nil {
		return errors.Wrap(err, "UserStore.UpdateGraduate got error")
	}
	return nil
}

func (*UserStore) UpdateEntered(db repository.DBHandler, studentNum int) error {

	if !exitsRecord(studentNum) {
		return errors.Errorf("no record with student_num = %d \n", studentNum)
	}

	q := `
		update users set entered = ? where student_num = ?
		`

	if _, err := db.Exec(q, true, studentNum); err != nil {
		return errors.Wrap(err, "UserStore.UpdateEntered got error")
	}

	return nil
}

func exitsRecord(studentNum int) bool {
	q := `
			select * from users where student_num = ? limit 1
		`
	var user model.User
	if err := db.Driver.QueryRowx(q, studentNum).StructScan(&user); err != nil {
		return false
	}
	return true
}
