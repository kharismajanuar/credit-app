package response

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
} //@name Response

type ResponseError struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
} //@name ResponseError

func ResponseSuccess(c *fiber.Ctx, statusCode int, message string, data any) error {
	response := Response{
		Status:  "Success",
		Message: message,
		Data:    data,
	}
	return c.Status(statusCode).JSON(response)
}

func ResponseFailed(c *fiber.Ctx, statusCode int, message any) error {
	response := ResponseError{
		Status:  "Failed",
		Message: message,
	}
	log.Printf("error: %v", message)
	return c.Status(statusCode).JSON(response)
}
