package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kharismajanuar/credit-app/internal/domain"
	"github.com/kharismajanuar/credit-app/internal/health/constant"
	"github.com/kharismajanuar/credit-app/utils/response"
)

type HealthHandler struct {
	heatlhService domain.HealthService
}

func NewHealthHandler(heatlhService domain.HealthService) *HealthHandler {
	return &HealthHandler{
		heatlhService: heatlhService,
	}
}

func (h *HealthHandler) Ping(c *fiber.Ctx) error {
	return response.ResponseSuccess(c, fiber.StatusOK, constant.Pong, nil)
}

func (h *HealthHandler) HealthCheck(c *fiber.Ctx) error {
	heatlh, err := h.heatlhService.HealthCheck()
	if err != nil {
		return response.ResponseFailed(c, fiber.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(c, fiber.StatusOK, constant.SuccessHealthCheck, formatHealthCheck(heatlh))
}
