package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ramadhan1445sprint/sprint_ecommerce/crypto"
	"github.com/ramadhan1445sprint/sprint_ecommerce/customErr"
)

func Authorization(ctx *fiber.Ctx) error {
	auth := ctx.Get("Authorization")
	if auth == "" {
		return customErr.NewUnauthorizedError("token not found")
	}

	splitted := strings.Split(auth, " ")

	if splitted[0] != "Bearer" {
		return customErr.NewUnauthorizedError("invalid token")
	}

	payload, err := crypto.VerifyToken(splitted[1])
	if err != nil {
		return customErr.NewUnauthorizedError(err.Error())
	}

	ctx.Locals("user_id", payload.Id)
	ctx.Locals("username", payload.Username)
	ctx.Locals("name", payload.Name)
	return ctx.Next()
}
