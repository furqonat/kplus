package routes

import (
	"log"

	"kplus.com/controllers"
	"kplus.com/middlewares"
	"kplus.com/utils"
)

type InstallmentRoute struct {
	controllers controllers.InstallmentController
	handler     utils.RequestHandler
	middlewares middlewares.JwtMiddleware
}

func (i InstallmentRoute) Setup() {

	log.Println("setting up installment routes")
	installment := i.handler.App.Group("/installment")
	installment.Post("/pay", i.controllers.PayInstallment)
}

func NewInstallmentRoute(
	controllers controllers.InstallmentController,
	handler utils.RequestHandler,
	jwt middlewares.JwtMiddleware,
) InstallmentRoute {
	return InstallmentRoute{controllers, handler, jwt}
}
