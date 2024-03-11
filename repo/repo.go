package repo

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type RepoInterface interface {
	GetStatus() (string, error)
}

func NewRepo(db *sqlx.DB) RepoInterface {
	return &repo{db: db}
}

type repo struct {
	db *sqlx.DB
}

type Status struct {
	Status string `db:"status"`
}

func (r *repo) GetStatus() (string, error) {
	var status []Status
	r.db.Select(&status, "SELECT * FROM status_check")

	if len(status) > 0 {
		return status[0].Status, nil
	}

	return "", errors.New("No Status Found")
}
