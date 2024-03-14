package svc

import (

	"github.com/ramadhan1445sprint/sprint_ecommerce/entity"
	"github.com/ramadhan1445sprint/sprint_ecommerce/repo"
	"github.com/ramadhan1445sprint/sprint_ecommerce/utils"
)

type StockSvcInterface interface {
	UpdateStock(productReq *entity.StockUpdateRequest, productId string) (int, error)
}

func NewStockSvc(repo repo.StockRepoInterface) StockSvcInterface {
	return &stockSvc{repo: repo}
}

type stockSvc struct {
	repo repo.StockRepoInterface
}

func (s *stockSvc) UpdateStock(productReq *entity.StockUpdateRequest, productId string) (int, error) {
	productReq.ProductID = &productId
	status, err := utils.ValidateStockUpdateRequest(productReq)

	if err != nil {
		return status, err
	}

	product := entity.Product{
		ID: *productReq.ProductID,
		Stock: *productReq.Stock,
	}

	err = s.repo.UpdateStock(&product)
	if err != nil {
		return 500, err
	}

	return  200, nil
}
