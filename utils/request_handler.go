package utils

import "github.com/gofiber/fiber/v2"

type RequestHandler struct {
	App *fiber.App
}

func NewRequestHandler() RequestHandler {
	app := fiber.New(fiber.Config{})
	return RequestHandler{
		App: app,
	}
}

type ResponseOk[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
	Total   *int64 `json:"total,omitempty"`
}

type ResponseError struct {
	Message string `json:"message"`
}
