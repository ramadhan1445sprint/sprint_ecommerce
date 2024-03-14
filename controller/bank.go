package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ramadhan1445sprint/sprint_ecommerce/entity"
	"github.com/ramadhan1445sprint/sprint_ecommerce/svc"
)

type BankAccountController struct {
	svc svc.BankAccountSvcInterface
}

func NewBankAccountController(svc svc.BankAccountSvcInterface) *BankAccountController {
	return &BankAccountController{svc: svc}
}

func (c *BankAccountController) GetBankAccount(ctx *fiber.Ctx) error {
	userId := 1
	res, status, err := c.svc.GetBankAccount(userId)

	if err != nil {
		ctx.Status(status).JSON(fiber.Map{"message": err.Error()})
	}

	resp := entity.Response{
		Message: "success",
		Data: res,
	}

	return ctx.Status(status).JSON(resp)
}

func (c *BankAccountController) CreateBankAccount(ctx *fiber.Ctx) error {
	var message string
	userId := 1

	req := new(entity.BankAccountCreateRequest)

	if err := ctx.BodyParser(req); err != nil {
		ctx.Status(500).JSON(fiber.Map{"message": "internal server error"})
	}

	status, err := c.svc.CreateBankAccount(req, userId)

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

func (c *BankAccountController) UpdateBankAccount(ctx *fiber.Ctx) error {
	var message string

	bankAccountId, err := ctx.ParamsInt("bankAccountId")

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{"message": "invalid bank id"})
	}

	req := new(entity.BankAccountUpdateRequest)

	if err := ctx.BodyParser(req); err != nil {
		ctx.Status(500).JSON(fiber.Map{"message": "internal server error"})
	}

	req.ID = &bankAccountId

	status, err := c.svc.UpdateBankAccount(req)

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

func (c *BankAccountController) DeleteBankAccount(ctx *fiber.Ctx) error {
	var message string

	bankAccountId, err := ctx.ParamsInt("bankAccountId")

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{"message": "invalid bank id"})
	}

	status, err := c.svc.DeleteBankAccount(bankAccountId)

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
