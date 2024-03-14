package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/ramadhan1445sprint/sprint_ecommerce/entity"
)

type UserRepo interface {
	GetUser(username string) (*entity.User, error)
	CreateUser(name, username, password string) error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) GetUser(username string) (*entity.User, error) {
	var user entity.User
	err := r.db.Get(&user, "SELECT id, name, username, password FROM users WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) CreateUser(name, username, password string) error {
	_, err := r.db.Exec("INSERT INTO users (name, username, password) VALUES ($1, $2, $3)", name, username, password)
	return err
}
