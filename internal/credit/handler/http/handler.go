package http

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/kharismajanuar/credit-app/internal/credit/model"
	"github.com/kharismajanuar/credit-app/internal/domain"
	sharedHelper "github.com/kharismajanuar/credit-app/internal/shared/helper"
	"github.com/kharismajanuar/credit-app/utils/response"
)

type CreditHandler struct {
	creditService domain.CreditService
}

func NewCreditHandler(creditService domain.CreditService) *CreditHandler {
	return &CreditHandler{
		creditService: creditService,
	}
}

// CreateCredit godoc
// @Summary Create Credit
// @Tags		Credit
// @Security	 Bearer
// @Accept		json
// @Produce	json
// @Param CreateCreditRequest body model.CreateCreditRequest true "Create Credit Request"
// @Success 201 {object} response.Response  "Response Success"
// @Failure 0 {object} response.ResponseError "Response Error"
// @Router		/credits	[post]
func (h *CreditHandler) CreateCredit(c *fiber.Ctx) error {
	var request model.CreateCreditRequest

	if err := c.BodyParser(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err.Error())
	}

	claim := sharedHelper.ExtractClaim(c)

	if claim.Role != "user" {
		return response.ResponseFailed(c, fiber.StatusUnauthorized, errors.New("only user can create credit").Error())
	}

	request.UserID = claim.UserID

	if err := sharedHelper.Validate(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err)
	}

	if err := h.creditService.CreateCredit(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.ResponseSuccess(c, fiber.StatusCreated, "Success create credit", nil)
}

// GetListCredit godoc
// @Summary Get List Credit
// @Tags		Credit
// @Security	 Bearer
// @Accept		json
// @Produce	json
// @Param user_id path int true "Get List Credit Request ID"
// @Success 200 {object} response.Response{data=[]ListCreditFormatter}  "Response Success"
// @Failure 0 {object} response.ResponseError "Response Error"
// @Router		/credits/{user_id}	[get]
func (h *CreditHandler) GetListCredit(c *fiber.Ctx) error {
	var request model.GetListCreditRequest

	if err := c.ParamsParser(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err.Error())
	}

	if err := sharedHelper.Validate(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err)
	}

	listCredit, err := h.creditService.GetListCredit(request.UserID)
	if err != nil {
		return response.ResponseFailed(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.ResponseSuccess(c, fiber.StatusOK, "Success get list credit", formatListCredit(listCredit))
}

// UpdateStatusCredit godoc
// @Summary Update Status Credit
// @Tags		Credit
// @Security	 Bearer
// @Accept		json
// @Produce	json
// @Param credit_id path int true "Update Status Credit Request ID"
// @Param UpdateStatusCreditRequest body model.UpdateStatusCreditRequest true "Update Status Credit Request"
// @Success 200 {object} response.Response  "Response Success"
// @Failure 0 {object} response.ResponseError "Response Error"
// @Router		/credits/{credit_id}	[put]
func (h *CreditHandler) UpdateStatusCredit(c *fiber.Ctx) error {
	var request model.UpdateStatusCreditRequest

	if err := c.ParamsParser(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err.Error())
	}

	if err := c.BodyParser(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err.Error())
	}

	claim := sharedHelper.ExtractClaim(c)
	if claim.Role != "manager" {
		return response.ResponseFailed(c, fiber.StatusUnauthorized, errors.New("only manager can update status credit").Error())
	}

	if err := sharedHelper.Validate(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err)
	}

	if err := h.creditService.UpdateStatusCredit(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.ResponseSuccess(c, fiber.StatusOK, "Success update status credit", nil)
}

// CreatePayment godoc
// @Summary Create Payment
// @Tags		Credit
// @Security	 Bearer
// @Accept		json
// @Produce	json
// @Param credit_id path int true "Create Payment Request ID"
// @Param CreatePaymentRequest body model.CreatePaymentRequest true "Credit Payment Request"
// @Success 201 {object} response.Response  "Response Success"
// @Failure 0 {object} response.ResponseError "Response Error"
// @Router		/credits/{credit_id}/payments	[post]
func (h *CreditHandler) CreatePayment(c *fiber.Ctx) error {
	var request model.CreatePaymentRequest

	if err := c.ParamsParser(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err.Error())
	}

	if err := c.BodyParser(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err.Error())
	}

	claim := sharedHelper.ExtractClaim(c)
	if claim.Role != "user" {
		return response.ResponseFailed(c, fiber.StatusUnauthorized, errors.New("only user can create payment").Error())
	}

	if err := sharedHelper.Validate(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err)
	}

	if err := h.creditService.CreatePayment(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.ResponseSuccess(c, fiber.StatusCreated, "Success create payment", nil)
}

// GetListPayment godoc
// @Summary Get List Payment
// @Tags		Credit
// @Security	 Bearer
// @Accept		json
// @Produce	json
// @Param credit_id path int true "Get List Payment Request ID"
// @Success 200 {object} response.Response{data=[]ListPaymentFormatter}  "Response Success"
// @Failure 0 {object} response.ResponseError "Response Error"
// @Router		/credits/{credit_id}/payments [get]
func (h *CreditHandler) GetListPayment(c *fiber.Ctx) error {
	var request model.GetListPaymentRequest

	if err := c.ParamsParser(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err.Error())
	}

	if err := sharedHelper.Validate(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err)
	}

	listPayment, err := h.creditService.GetListPayment(request.CreditID)
	if err != nil {
		return response.ResponseFailed(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.ResponseSuccess(c, fiber.StatusOK, "Success get list payment", formatListPayment(listPayment))
}

// UpdateStatusPayment godoc
// @Summary Update Status Payment
// @Tags		Credit
// @Security	 Bearer
// @Accept		json
// @Produce	json
// @Param payment_id path int true "Update Status Payment Request ID"
// @Param UpdateStatusPaymentRequest body model.UpdateStatusPaymentRequest true "Update Status Payment Request"
// @Success 200 {object} response.Response  "Response Success"
// @Failure 0 {object} response.ResponseError "Response Error"
// @Router		/payments/{payment_id}	[put]
func (h *CreditHandler) UpdateStatusPayment(c *fiber.Ctx) error {
	var request model.UpdateStatusPaymentRequest

	if err := c.ParamsParser(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err.Error())
	}

	if err := c.BodyParser(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err.Error())
	}

	claim := sharedHelper.ExtractClaim(c)
	if claim.Role != "admin" {
		return response.ResponseFailed(c, fiber.StatusUnauthorized, errors.New("only admin can update status payment").Error())
	}

	if err := sharedHelper.Validate(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err)
	}

	if err := h.creditService.UpdateStatusPayment(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.ResponseSuccess(c, fiber.StatusOK, "Success update status payment", nil)
}
