package routes

import (
	"kplus.com/controllers"
	"kplus.com/middlewares"
	"kplus.com/utils"
)

type TransactionRoute struct {
	controller controllers.TransactionController
	middleware middlewares.JwtMiddleware
	handler    utils.RequestHandler
}

func (t TransactionRoute) Setup() {
	api := t.handler.App.Group("/trx")
	api.Post("/", t.middleware.HandleAuthWithRoles(
		utils.RoleUser,
	), t.controller.CreateTransaction)
	api.Get("/:id", t.middleware.HandleAuthWithRoles(
		utils.RoleUser,
	), t.controller.GetTransaction)
	api.Get("/", t.middleware.HandleAuthWithRoles(
		utils.RoleUser,
	), t.controller.GetTransactions)
}

func NewTransactionRoute(
	controller controllers.TransactionController,
	middleware middlewares.JwtMiddleware,
	handler utils.RequestHandler) TransactionRoute {
	return TransactionRoute{
		controller: controller,
		middleware: middleware,
		handler:    handler,
	}
}
