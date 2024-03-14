package svc

import (
	"errors"

	"github.com/ramadhan1445sprint/sprint_ecommerce/entity"
	"github.com/ramadhan1445sprint/sprint_ecommerce/repo"
	"github.com/ramadhan1445sprint/sprint_ecommerce/utils"
)

type PaymentSvcInterface interface {
	CreatePayment(paymentReq *entity.PaymentCreateRequest) (int, error)
}

func NewPaymentSvc(repo repo.PaymentRepoInterface) PaymentSvcInterface {
	return &paymentSvc{repo: repo}
}

type paymentSvc struct {
	repo repo.PaymentRepoInterface
}

func (s *paymentSvc) CreatePayment(paymentReq *entity.PaymentCreateRequest) (int, error) {
	status, err := utils.ValidatePaymentRequest(paymentReq)

	if err != nil {
		return status, err
	}

	productQty := 20

	if productQty < *paymentReq.Quantity {
		return 400, errors.New("insufficient quantity")
	}

	payment := entity.Payment{
		ProductID: *paymentReq.ProductID,
		BankAccountID: *paymentReq.BankAccountID,
		PaymentProofImgUrl: *paymentReq.PaymentProofImgUrl,
		Quantity: *paymentReq.Quantity,
	}

	err = s.repo.CreatePayment(&payment)
	if err != nil {
		return 500, err
	}

	return  200, nil
}
