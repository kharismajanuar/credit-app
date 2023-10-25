package http

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kharismajanuar/credit-app/internal/domain"
	sharedConstant "github.com/kharismajanuar/credit-app/internal/shared/constant"
	sharedHelper "github.com/kharismajanuar/credit-app/internal/shared/helper"
	"github.com/kharismajanuar/credit-app/internal/user/model"
	"github.com/kharismajanuar/credit-app/utils/response"
)

type UserHandler struct {
	userService domain.UserService
}

func NewUserHandler(userService domain.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetListUser godoc
// @Summary Get List User
// @Tags		User
// @Security Bearer
// @Accept		json
// @Produce	json
// @Success 201 {object} response.Response{data=[]ListUserFormatter}  "Response Success"
// @Failure 0 {object} response.ResponseError "Response Error"
// @Router		/users [get]
func (h *UserHandler) GetListUser(c *fiber.Ctx) error {
	listUser, err := h.userService.GetListUser()
	if err != nil {
		return response.ResponseFailed(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.ResponseSuccess(c, fiber.StatusOK, "Success get list user", formatListUser(listUser))
}

// GetDetailUser godoc
// @Summary Get Detail User
// @Tags		User
// @Security Bearer
// @Accept		json
// @Produce	json
// @Param user_id path int true "Get Detail User Request"
// @Success 201 {object} response.Response{data=DetailUserFormatter}  "Response Success"
// @Failure 0 {object} response.ResponseError "Response Error"
// @Router		/users/{user_id} [get]
func (h *UserHandler) GetDetailUser(c *fiber.Ctx) error {
	var request model.GetDetailUserRequest

	if err := c.ParamsParser(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err.Error())
	}

	if err := sharedHelper.Validate(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err)
	}

	user, err := h.userService.GetDetailUser(request.UserID)
	if err != nil {
		return response.ResponseFailed(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.ResponseSuccess(c, fiber.StatusOK, "Success get detail user", formatDetailUser(user))
}

// UpdateUser godoc
// @Summary Update User
// @Tags		User
// @Security Bearer
// @Accept		json
// @Produce	json
// @Param user_id path int true "Update User Request ID"
// @Param UpdateUserRequest body model.UpdateUserRequest true "Update User Request Body"
// @Success 200 {object} response.Response  "Response Success"
// @Failure 0 {object} response.ResponseError "Response Error"
// @Router		/users/{user_id} [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var request model.UpdateUserRequest

	if err := c.ParamsParser(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err.Error())
	}

	if err := c.BodyParser(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err.Error())
	}

	// Validate user
	claim := sharedHelper.ExtractClaim(c)
	if claim.UserID != request.UserID {
		return response.ResponseFailed(c, fiber.StatusUnauthorized, errors.New(sharedConstant.AccessDenied).Error())
	}

	if err := sharedHelper.Validate(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err)
	}

	if !sharedHelper.ValidateDateFormat(request.DateOfBirth) {
		err := fmt.Errorf(sharedConstant.InvalidDateLayout, request.DateOfBirth)
		return response.ResponseFailed(c, fiber.StatusBadRequest, err.Error())
	}

	if err := h.userService.UpdateUser(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.ResponseSuccess(c, fiber.StatusOK, "Success update user", nil)
}

// DeleteUser godoc
// @Summary Delete User
// @Tags		User
// @Security Bearer
// @Accept		json
// @Produce	json
// @Param user_id path int true "Delete User Request ID"
// @Success 200 {object} response.Response  "Response Success"
// @Failure 0 {object} response.ResponseError "Response Error"
// @Router		/users/{user_id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	var request model.DeleteUserRequest

	if err := c.ParamsParser(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err.Error())
	}

	// Validate user
	claim := sharedHelper.ExtractClaim(c)
	if claim.UserID != request.UserID {
		return response.ResponseFailed(c, fiber.StatusUnauthorized, errors.New(sharedConstant.AccessDenied).Error())
	}

	if err := sharedHelper.Validate(&request); err != nil {
		return response.ResponseFailed(c, fiber.StatusBadRequest, err)
	}

	if err := h.userService.DeleteUser(request.UserID); err != nil {
		return response.ResponseFailed(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.ResponseSuccess(c, fiber.StatusOK, "Success delete user", nil)
}
