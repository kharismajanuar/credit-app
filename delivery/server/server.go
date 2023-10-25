package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Http(handler HandlerHttp) *fiber.App {
	app := fiber.New()

	// Init global middlewares
	app.Use(cors.New())
	app.Use(favicon.New())
	app.Use(logger.New(logger.Config{
		Format:     "${time} [${ip}]:${port} ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006/01/02 15:04:05",
	}))
	app.Use(recover.New())

	// Init router
	router(app, handler)

	// Prepare an endpoint for 'Not Found'
	app.Use(func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})

	return app
}
