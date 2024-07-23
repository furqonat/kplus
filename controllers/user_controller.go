package controllers

import (
	"github.com/gofiber/fiber/v2"
	"kplus.com/dto"
	"kplus.com/services"
	"kplus.com/utils"
)

type UserController struct {
	service services.UserService
}

func (u UserController) GetUser(c *fiber.Ctx) error {
	token := c.Locals(utils.Token).(utils.JwtCustomClaims)
	response, err := u.service.GetUser(token.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError{
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseOk[*dto.UserDto]{
		Data:    response,
		Message: "get user success",
	})
}

func (u UserController) CreateUserDetails(c *fiber.Ctx) error {
	body := dto.UserDetailsDto{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError{
			Message: err.Error(),
		})
	}
	response, err := u.service.CreateUserDetails(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError{
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseOk[int64]{
		Data:    response,
		Message: "create user details success",
	})
}

func (u UserController) UpdateUserDetails(c *fiber.Ctx) error {
	body := dto.UserDetailsDto{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError{
			Message: err.Error(),
		})
	}
	token := c.Locals(utils.Token).(utils.JwtCustomClaims)
	if err := u.service.UpdateUserDetails(body, token.UserID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError{
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseOk[int64]{
		Message: "update user details success",
	})
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		service: userService,
	}
}
