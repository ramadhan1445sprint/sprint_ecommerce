package controller

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ramadhan1445sprint/sprint_ecommerce/repo"
	"github.com/ramadhan1445sprint/sprint_ecommerce/svc"
	"github.com/ramadhan1445sprint/sprint_ecommerce/utils"
)

type Controller struct {
	svc svc.SvcInterface
}

func NewController(svc svc.SvcInterface) *Controller {
	return &Controller{svc: svc}
}

func ValidateCondition(fl validator.FieldLevel) bool {
	condition := fl.Field().String()
	return condition == "new" || condition == "second"
}

func (c *Controller) CreateProduct(ctx *fiber.Ctx) error {
	// cek authenthication

	product := &repo.Product{}

	// Check, if received JSON data is valid.
	if err := ctx.BodyParser(product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	product.ID = uuid.New()
	product.CreatedAt = time.Now().UTC()
	product.UpdatedAt = time.Now().UTC()

	validate := validator.New()

	validate.RegisterValidation("validCondition", ValidateCondition)

	// Validate book fields.
	if err := validate.Struct(product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": utils.ValidatorErrors(err),
		})
	}

	err := c.svc.CreateProduct(*product)
	if err != nil {
		return ctx.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "product added successfully",
	})
}

func (c *Controller) UpdateProduct(ctx *fiber.Ctx) error {
	// cek authenthication

	product := &repo.Product{}

	id, err := uuid.Parse(ctx.Params("productId"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// check productId is found or not
	*product, err = c.svc.GetDetailProduct(id)
	if err != nil || product.ID == uuid.Nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "product with the given ID is not found",
		})
	}

	// Check, if received JSON data is valid.
	if err := ctx.BodyParser(product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	product.ID = id
	product.UpdatedAt = time.Now().UTC()

	validate := validator.New()

	validate.RegisterValidation("validCondition", ValidateCondition)

	// Validate book fields.
	if err := validate.Struct(product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": utils.ValidatorErrors(err),
		})
	}

	err = c.svc.UpdateProduct(*product)
	if err != nil {
		return ctx.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "product updated successfully",
	})
}

func (c *Controller) GetDetailProduct(ctx *fiber.Ctx) error {
	// cek authenthication

	id, err := uuid.Parse(ctx.Params("productId"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	product, err := c.svc.GetDetailProduct(id)
	if err != nil || product.ID == uuid.Nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "product with the given ID is not found",
		})
	}

	// Return status 200 OK.
	return ctx.JSON(fiber.Map{
		"message": "ok",
		"data": fiber.Map{
			"product": fiber.Map{
				"productId":      product.ID,
				"name":           product.Name,
				"price":          product.Price,
				"imageUrl":       product.ImageUrl,
				"stock":          product.Stock,
				"condition":      product.Condition,
				"tags":           product.Tags,
				"isPurchaseable": product.IsPurchasable,
			},
			"seller": "",
		},
	})
}

func (c *Controller) DeleteProduct(ctx *fiber.Ctx) error {
	// cek authenthication

	product := &repo.Product{}

	id, err := uuid.Parse(ctx.Params("productId"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// check productId is found or not
	*product, err = c.svc.GetDetailProduct(id)
	fmt.Println(product)
	if err != nil || product.ID == uuid.Nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "product with the given ID is not found",
		})
	}

	err = c.svc.DeleteProduct(id)
	if err != nil {
		return ctx.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "product deleted successfully",
	})
}

func (c *Controller) GetListProduct(ctx *fiber.Ctx) error {
	// cek authenthication

	// Extract keys from the query parameters map
	keys := &repo.Key{}

	var products []repo.Product
	var limit, offset int = 0, 0

	// Check, if received JSON data is valid.
	if err := ctx.QueryParser(keys); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	products, err := c.svc.GetListProduct(*keys)
	if err != nil {
		return ctx.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if keys.Limit != nil && keys.Offset != nil {
		limit = *keys.Limit
		offset = *keys.Offset
	}

	// Return status 200 OK.
	return ctx.JSON(fiber.Map{
		"message": "ok",
		"data":    products,
		"meta": fiber.Map{
			"limit":  limit,
			"offset": offset,
		},
	})
}
