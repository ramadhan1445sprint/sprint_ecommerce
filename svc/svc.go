package svc

import "github.com/ramadhan1445sprint/sprint_ecommerce/repo"

type SvcInterface interface {
	GetStatus() (string, error)
}

func NewSvc(repo repo.RepoInterface) SvcInterface {
	return &svc{repo: repo}
}

type svc struct {
	repo repo.RepoInterface
}

func (s *svc) GetStatus() (string, error) {
	status, err := s.repo.GetStatus()
	if err != nil {
		return "", err
	}

	return status, nil
}
