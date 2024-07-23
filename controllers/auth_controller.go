package controllers

import (
	"github.com/gofiber/fiber/v2"
	"kplus.com/dto"
	"kplus.com/services"
	"kplus.com/utils"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return AuthController{
		service: authService,
	}
}

func (a AuthController) SignIn(c *fiber.Ctx) error {
	body := dto.SignInDto{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError{
			Message: err.Error(),
		})
	}
	if response, err := a.service.SignIn(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError{
			Message: err.Error(),
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(utils.ResponseOk[*dto.ResponseSignInDto]{
			Data:    response,
			Message: "sign in success",
		})
	}
}

func (a AuthController) SignUp(c *fiber.Ctx) error {
	body := dto.SignUpDto{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError{
			Message: err.Error(),
		})
	}
	if response, err := a.service.SignUp(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError{
			Message: err.Error(),
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(utils.ResponseOk[*dto.ResponseSignInDto]{
			Data:    response,
			Message: "sign up success",
		})
	}
}

func (a AuthController) RefreshToken(c *fiber.Ctx) error {
	token := c.Locals(utils.Token).(utils.JwtCustomClaims)
	if response, err := a.service.RefreshToken(token); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError{
			Message: err.Error(),
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(utils.ResponseOk[*string]{
			Data:    response,
			Message: "refresh token success",
		})
	}
}
