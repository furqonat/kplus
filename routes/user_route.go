package routes

import (
	"log"

	"kplus.com/controllers"
	"kplus.com/middlewares"
	"kplus.com/utils"
)

type UserRoute struct {
	controllers controllers.UserController
	handler     utils.RequestHandler
	middlewares middlewares.JwtMiddleware
}

func (u UserRoute) Setup() {
	log.Println("setting up user routes")
	user := u.handler.App.Group("/user")
	user.Post("/", u.middlewares.HandleAuthWithRoles(
		utils.RoleUser,
	), u.controllers.CreateUserDetails)
	user.Put("/", u.middlewares.HandleAuthWithRoles(
		utils.RoleUser,
	), u.controllers.UpdateUserDetails)
	user.Get("/", u.middlewares.HandleAuthWithRoles(
		utils.RoleUser,
	), u.controllers.GetUser)
	user.Get("/loans", u.middlewares.HandleAuthWithRoles(
		utils.RoleUser,
	), u.controllers.GetLoanLimit)
}

func NewUserRoute(
	controllers controllers.UserController,
	handler utils.RequestHandler,
	jwt middlewares.JwtMiddleware,
) UserRoute {
	return UserRoute{controllers, handler, jwt}
}
