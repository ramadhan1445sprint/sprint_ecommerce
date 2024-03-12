package svc

import (
	"github.com/google/uuid"
	"github.com/ramadhan1445sprint/sprint_ecommerce/repo"
)

type SvcInterface interface {
	CreateProduct(product repo.Product) error
	GetDetailProduct(id uuid.UUID) (repo.Product, error)
	UpdateProduct(product repo.Product) error
	DeleteProduct(id uuid.UUID) error
}

func NewSvc(repo repo.RepoInterface) SvcInterface {
	return &svc{repo: repo}
}

type svc struct {
	repo repo.RepoInterface
}

func (s *svc) CreateProduct(product repo.Product) error {
	err := s.repo.CreateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (s *svc) GetDetailProduct(id uuid.UUID) (repo.Product, error) {
	product, err := s.repo.GetDetailProduct(id)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *svc) UpdateProduct(product repo.Product) error {
	err := s.repo.UpdateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (s *svc) DeleteProduct(id uuid.UUID) error {
	err := s.repo.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}
