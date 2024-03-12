package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/ramadhan1445sprint/sprint_ecommerce/controller"
	"github.com/ramadhan1445sprint/sprint_ecommerce/repo"
	"github.com/ramadhan1445sprint/sprint_ecommerce/svc"
)

type Server struct {
	db  *sqlx.DB
	app *fiber.App
}

func NewServer(db *sqlx.DB) *Server {
	app := fiber.New()

	return &Server{
		app: app,
		db:  db,
	}
}

func (s *Server) Run() {
	s.app.Listen(":8000")
}

func (s *Server) RegisterRoute() {
	mainRoute := s.app.Group("/v1")
	registerProducRouter(mainRoute, s.db)
}

func registerProducRouter(r fiber.Router, db *sqlx.DB) {
	c := controller.NewController(svc.NewSvc(repo.NewRepo(db)))
	prodRouter := r.Group("/product")

	// prodRouter.Get("/", c.Get)
	prodRouter.Get("/:productId", c.GetDetailProduct)
	prodRouter.Post("/", c.CreateProduct)
	prodRouter.Patch("/:productId", c.UpdateProduct)
	prodRouter.Delete("/:productId", c.DeleteProduct)
}
