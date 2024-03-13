package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ramadhan1445sprint/sprint_ecommerce/err"
)

func ErrorHandler(ctx *fiber.Ctx, rawErr error) error {
	code := fiber.StatusInternalServerError
	if customErr, ok := rawErr.(err.CustomError); ok {
		code = customErr.Status()
		return ctx.Status(code).JSON(fiber.Map{
			"message": customErr.Error(),
		})
	}

	return ctx.Status(code).JSON(fiber.Map{
		"message": rawErr.Error(),
	})
}
