package datastore

import (
	"github.com/pkg/errors"
	"github.com/shinnosuke-K/lab-assignment-backend/domain/model"
	"github.com/shinnosuke-K/lab-assignment-backend/domain/repository"
)

type AnswerStore struct{}

func (*AnswerStore) Store(db repository.DBHandler, answers []*model.Answer) error {
	return nil
}

func (*AnswerStore) Fix(db repository.DBHandler, correctedAns []*model.Answer) error {
	return nil
}

func (*AnswerStore) GetAnswersByUserID(db repository.DBHandler, userID string) ([]*model.Answer, error) {

	q := `
			select * from answers where user_id = ?
		`

	var res []*model.Answer
	if err := db.QueryRowx(q, userID).StructScan(&res); err != nil {
		return nil, errors.Wrap(err, "AnswerStore.GetAnswersByUserID go error")
	}
	return res, nil
}
