package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/kharismajanuar/credit-app/delivery/server/middleware"
	_ "github.com/kharismajanuar/credit-app/docs"
)

func router(app *fiber.App, handler HandlerHttp) {

	jwt := middleware.NewAuthMiddleware()

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/ping", handler.healthHandler.Ping)
	app.Get("/health-check", handler.healthHandler.HealthCheck)

	app.Post("/register", handler.authHandler.Register)
	app.Post("/login", handler.authHandler.Login)

	user := app.Group("/users")
	{
		user.Get("/", jwt, handler.userHandler.GetListUser)
		user.Get("/:user_id", jwt, handler.userHandler.GetDetailUser)
		user.Put("/:user_id", jwt, handler.userHandler.UpdateUser)
		user.Delete("/:user_id", jwt, handler.userHandler.DeleteUser)
	}

	credit := app.Group("/credits")
	{
		credit.Post("/", jwt, handler.creditHandler.CreateCredit)
		credit.Get("/:user_id", jwt, handler.creditHandler.GetListCredit)
		credit.Put("/:credit_id", jwt, handler.creditHandler.UpdateStatusCredit)
		credit.Post("/:credit_id/payments", jwt, handler.creditHandler.CreatePayment)
		credit.Get("/:credit_id/payments", jwt, handler.creditHandler.GetListPayment)
	}

	payment := app.Group("/payments")
	{
		payment.Put("/:payment_id", jwt, handler.creditHandler.UpdateStatusPayment)
	}

}
