package model

import "time"

type Answer struct {
	ID       int       `db:"id"`
	UserID   string    `db:"user_id"`
	ProfID   string    `db:"prof_id"`
	Point    int       `db:"point"`
	RegAt    time.Time `db:"reg_at"`
	UpdateAt time.Time `db:"update_at"`
}
