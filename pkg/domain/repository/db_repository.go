package repository

import "github.com/jmoiron/sqlx"

type DBHandler interface {
	sqlx.Execer
	sqlx.Preparer
	sqlx.Queryer
}

type DBConnector struct {
	*sqlx.DB
}
