package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ramadhan1445sprint/sprint_ecommerce/svc"
)

type Controller struct {
	svc svc.SvcInterface
}

func NewController(svc svc.SvcInterface) *Controller {
	return &Controller{svc: svc}
}

func (c *Controller) Get(ctx *fiber.Ctx) error {
	status, err := c.svc.GetStatus()
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"status": status,
	})
}