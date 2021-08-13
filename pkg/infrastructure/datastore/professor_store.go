package datastore

import (
	"github.com/pkg/errors"
	"github.com/shinnosuke-K/lab-assignment-backend/pkg/domain/model"
	"github.com/shinnosuke-K/lab-assignment-backend/pkg/domain/repository"
)

type ProfStore struct{}

func (*ProfStore) GetAll(db repository.DBHandler) ([]*model.Professor, error) {

	q := `
			select * from professors
		`

	var res []*model.Professor
	if err := db.QueryRowx(q).StructScan(&res); err != nil {
		return nil, errors.Wrap(err, "ProfStore.GetAll go error")
	}
	return res, nil
}

func (*ProfStore) GetByID(db repository.DBHandler, id string) (*model.Professor, error) {

	q := `
			select * from professors where id = ?
		`

	var res *model.Professor
	if err := db.QueryRowx(q, id).StructScan(&res); err != nil {
		return nil, errors.Wrap(err, "ProfStore.GetByID go errror")
	}
	return res, nil
}
