package controllers

import (
	"github.com/gofiber/fiber/v2"
	"kplus.com/dto"
	"kplus.com/services"
	"kplus.com/utils"
)

type TransactionController struct {
	transactionService services.TransactionService
}

func (t TransactionController) GetTransactions(c *fiber.Ctx) error {
	token := c.Locals(utils.Token).(utils.JwtCustomClaims)
	userID := utils.StringToInt(token.UserID)
	data, err := t.transactionService.GetTransactions(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError{
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseOk[[]dto.TransactionDto]{
		Data:    data,
		Message: "get transactions success",
	})
}

func (t TransactionController) GetTransaction(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := t.transactionService.GetTransaction(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError{
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseOk[*dto.TransactionDto]{
		Data:    data,
		Message: "get transaction success",
	})
}

func (t TransactionController) CreateTransaction(c *fiber.Ctx) error {
	body := dto.CreateTransactionDto{}

	token := c.Locals(utils.Token).(utils.JwtCustomClaims)
	userID := utils.StringToInt(token.UserID)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError{
			Message: err.Error(),
		})
	}
	err := t.transactionService.CreateTransaction(&body, userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError{
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseOk[string]{
		Data:    "create transaction success",
		Message: "create transaction success",
	})
}

func NewTransactionController(db utils.Database) TransactionController {
	return TransactionController{
		transactionService: services.NewTransactionService(db),
	}
}
