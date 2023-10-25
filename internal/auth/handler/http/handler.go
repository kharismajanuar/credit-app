package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kharismajanuar/credit-app/internal/auth/model"
	"github.com/kharismajanuar/credit-app/internal/domain"
	sharedConstant "github.com/kharismajanuar/credit-app/internal/shared/constant"
	sharedHelper "github.com/kharismajanuar/credit-app/internal/shared/helper"
	"github.com/kharismajanuar/credit-app/utils/response"
)

type AuthHandler struct {
	authService domain.AuthService
}

func NewAuthHandler(authService domain.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register godoc
// @Summary Register
// @Tags		Auth
// @Accept		json
// @Produce	json
// @Param RegisterRequest body model.RegisterRequest true "Register Request"
// @Success 201 {object} response.Response  "Response Success"
// @Failure 0 {object} response.ResponseError "Response Error"
// @Router		/register	[post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var request model.RegisterRequest

	if err := c.BodyParser(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err.Error())
	}

	if err := sharedHelper.Validate(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err)
	}

	if !sharedHelper.ValidateAlphanumericPunct(request.Password) {
		return response.ResponseFailed(c, fiber.StatusBadRequest, sharedConstant.InvalidPasswordInput)
	}

	if err := h.authService.Register(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.ResponseSuccess(c, fiber.StatusCreated, "Success register", nil)
}

// Login godoc
// @Summary Login
// @Tags		Auth
// @Accept		json
// @Produce	json
// @Param LoginRequest body model.LoginRequest true "Login Request"
// @Success 201 {object} response.Response{data=http.LoginFormatter}  "Response Success"
// @Failure 0 {object} response.ResponseError "Response Error"
// @Router		/login	[post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var request model.LoginRequest

	if err := c.BodyParser(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err.Error())
	}

	if err := sharedHelper.Validate(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err)
	}

	if !sharedHelper.ValidateAlphanumericPunct(request.Password) {
		return response.ResponseFailed(c, fiber.StatusBadRequest, sharedConstant.InvalidPasswordInput)
	}

	auth, err := h.authService.Login(&request)
	if err != nil {
		return response.ResponseFailed(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.ResponseSuccess(c, fiber.StatusOK, "Success login", formatLogin(auth))
}
