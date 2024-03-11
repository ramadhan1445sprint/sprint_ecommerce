package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ramadhan1445sprint/sprint_ecommerce/crypto"
)

func Authorization(ctx *fiber.Ctx) error {
	auth := ctx.Get("Authorization")
	splitted := strings.Split(auth, " ")

	if splitted[0] != "Bearer" {
		return ctx.JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	payload, err := crypto.VerifyToken(splitted[1])
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	ctx.Locals("username", payload.Username)
	ctx.Locals("name", payload.Name)
	return ctx.Next()
}
