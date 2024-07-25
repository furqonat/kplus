package controllers

import (
	"github.com/gofiber/fiber/v2"
	"kplus.com/dto"
	"kplus.com/services"
	"kplus.com/utils"
)

type InstallmentController struct {
	service services.InstallmentService
}

func (i InstallmentController) PayInstallment(c *fiber.Ctx) error {
	data := dto.PayInstallmentDto{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError{
			Message: err.Error(),
		})
	}

	if err := i.service.PayInstallment(data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError{
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseOk[string]{
		Message: "success",
	})
}

func NewInstallmentController(service services.InstallmentService) InstallmentController {
	return InstallmentController{
		service: service,
	}
}
