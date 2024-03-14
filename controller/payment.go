package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ramadhan1445sprint/sprint_ecommerce/entity"
	"github.com/ramadhan1445sprint/sprint_ecommerce/svc"
)

type paymentController struct {
	svc svc.PaymentSvcInterface
}

func NewPaymentController(svc svc.PaymentSvcInterface) *paymentController {
	return &paymentController{svc: svc}
}

func (c *paymentController) CreatePayment(ctx *fiber.Ctx) error {
	var message string
	
	productId, err := ctx.ParamsInt("productId")

	if err != nil {
		ctx.Status(400).JSON(fiber.Map{"message": "invalid product id"})
	}

	req := new(entity.PaymentCreateRequest)

	if err := ctx.BodyParser(req); err != nil {
		ctx.Status(500).JSON(fiber.Map{"message": "internal server error"})
	}

	req.ProductID = &productId

	status, err := c.svc.CreatePayment(req)

	if err != nil {
		message = err.Error()
	}else {
		message = "success"
	}

	resp := entity.Response{
		Message: message,
	}

	return ctx.Status(status).JSON(resp)
}
