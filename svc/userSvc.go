package svc

import (
	"errors"

	"github.com/ramadhan1445sprint/sprint_ecommerce/crypto"
	"github.com/ramadhan1445sprint/sprint_ecommerce/entity"
	"github.com/ramadhan1445sprint/sprint_ecommerce/repo"
)

type UserSvc interface {
	RegisterUser(*entity.User) (string, error)
	Login(entity.Credential) (*entity.User, string, error)
}

type userSvc struct {
	repo repo.UserRepo
}

func NewUserSvc(repo repo.UserRepo) UserSvc {
	return &userSvc{repo}
}

func (s *userSvc) RegisterUser(user *entity.User) (string, error) {
	existingUser, err := s.repo.GetUser(user.Username)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return "", err
		}
	}
	if existingUser != nil {
		return "", errors.New("username already exists")
	}

	// bcrypt user password
	hashedPassword, err := crypto.GenerateHashedPassword(user.Password)
	if err != nil {
		return "", err
	}

	err = s.repo.CreateUser(user.Name, user.Username, hashedPassword)
	if err != nil {
		return "", err
	}

	token, err := crypto.GenerateToken(user.Username, user.Name)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userSvc) Login(creds entity.Credential) (*entity.User, string, error) {
	user, err := s.repo.GetUser(creds.Username)
	if err != nil {
		return nil, "", err
	}

	err = crypto.VerifyPassword(creds.Password, user.Password)
	if err != nil {
		return nil, "", err
	}

	token, err := crypto.GenerateToken(user.Username, user.Name)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
