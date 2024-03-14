package server

import (
	"context"
	"log"

	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/ramadhan1445sprint/sprint_ecommerce/config"
	"github.com/ramadhan1445sprint/sprint_ecommerce/controller"
	"github.com/ramadhan1445sprint/sprint_ecommerce/middleware"
	"github.com/ramadhan1445sprint/sprint_ecommerce/repo"
	"github.com/ramadhan1445sprint/sprint_ecommerce/svc"
)

type Server struct {
	db  *sqlx.DB
	app *fiber.App
}

func NewServer(db *sqlx.DB) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

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
	registerUserRouter(mainRoute, s.db)
	healthCheckRoute(mainRoute, s.db)
	mainRoute.Use(middleware.Authorization)
	// put another route below
	registerImageRouter(mainRoute)
	registerProducRouter(mainRoute, s.db)
	productRoute(mainRoute, s.db)
	bankAccountRoute(mainRoute, s.db)
}

func registerProducRouter(r fiber.Router, db *sqlx.DB) {
	c := controller.NewController(svc.NewSvc(repo.NewRepo(db)))
	prodRouter := r.Group("/product")

	prodRouter.Get("/", c.GetListProduct)
	prodRouter.Get("/:productId", c.GetDetailProduct)
	prodRouter.Post("/", c.CreateProduct)
	prodRouter.Patch("/:productId", c.UpdateProduct)
	prodRouter.Delete("/:productId", c.DeleteProduct)
}
func registerUserRouter(r fiber.Router, db *sqlx.DB) {
	ctr := controller.NewUserController(svc.NewUserSvc(repo.NewUserRepo(db)))
	userGroup := r.Group("/user")

	userGroup.Post("/register", ctr.Register)
	userGroup.Post("/login", ctr.Login)
}

func registerImageRouter(r fiber.Router) {
	cfg, err := awsCfg.LoadDefaultConfig(
		context.Background(),
		awsCfg.WithRegion("ap-southeast-1"),
		awsCfg.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				config.GetString("S3_ID"),
				config.GetString("S3_SECRET_KEY"),
				"",
			),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	ctr := controller.NewImageController(svc.NewImageSvc(cfg))

	r.Post("/image", ctr.UploadImage)
}

func healthCheckRoute(r fiber.Router, db *sqlx.DB) {
	c := controller.NewHealthController(svc.NewHealthSvc(repo.NewHealthRepo(db)))
	healthRouter := r.Group("/health")
	healthRouter.Get("/", c.Get)
}

func productRoute(r fiber.Router, db *sqlx.DB) {
	payment := controller.NewPaymentController(svc.NewPaymentSvc(repo.NewPaymentRepo(db)))
	stock := controller.NewStockController(svc.NewStockSvc(repo.NewStockRepo(db)))
	productRouter := r.Group("/product")
	productRouter.Post("/:productId/buy", payment.CreatePayment)
	productRouter.Post("/:productId/stock", stock.UpdateStock)
}

func bankAccountRoute(r fiber.Router, db *sqlx.DB) {
	bankAccount := controller.NewBankAccountController(svc.NewBankAccounthSvc(repo.NewBankAccountRepo(db)))
	productRouter := r.Group("/bank")
	productRouter.Post("/account", bankAccount.CreateBankAccount)
	productRouter.Get("/account", bankAccount.GetBankAccount)
	productRouter.Patch("/account/:bankAccountId", bankAccount.UpdateBankAccount)
	productRouter.Delete("/account/:bankAccountId", bankAccount.DeleteBankAccount)
}

