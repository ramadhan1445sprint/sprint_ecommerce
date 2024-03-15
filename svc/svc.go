package svc

import (
	"github.com/google/uuid"
	"github.com/ramadhan1445sprint/sprint_ecommerce/entity"
	"github.com/ramadhan1445sprint/sprint_ecommerce/repo"
)

type SvcInterface interface {
	CreateProduct(product entity.Product) error
	GetDetailProduct(id uuid.UUID) (entity.Product, error)
	UpdateProduct(product entity.Product) error
	DeleteProduct(id uuid.UUID) error
	GetListProduct(keys entity.Key, userId uuid.UUID) ([]entity.Product, error)
	GetPurchaseCount(id uuid.UUID) (int, error)
	GetProductSoldTotal(userId uuid.UUID) (entity.ProductPayment, error)
	GetCountProduct() (int, error)
}

func NewSvc(repo repo.RepoInterface) SvcInterface {
	return &svc{repo: repo}
}

type svc struct {
	repo repo.RepoInterface
}

func (s *svc) CreateProduct(product entity.Product) error {
	err := s.repo.CreateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (s *svc) GetDetailProduct(id uuid.UUID) (entity.Product, error) {
	product, err := s.repo.GetDetailProduct(id)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *svc) UpdateProduct(product entity.Product) error {
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

func (s *svc) GetListProduct(keys entity.Key, userId uuid.UUID) ([]entity.Product, error) {
	product, err := s.repo.GetListProduct(keys, userId)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *svc) GetPurchaseCount(id uuid.UUID) (int, error) {
	total, err := s.repo.GetPurchaseCount(id)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (s *svc) GetProductSoldTotal(userId uuid.UUID) (entity.ProductPayment, error) {
	productPayment, err := s.repo.GetProductSoldTotal(userId)
	if err != nil {
		return productPayment, err
	}

	return productPayment, nil
}

func (s *svc) GetCountProduct() (int, error) {
	count, err := s.repo.GetCountProduct()
	if err != nil {
		return 0, err
	}

	return count, nil
}