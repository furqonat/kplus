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
	user.Post("/details", u.middlewares.HandleAuthWithRoles(
		utils.RoleUser,
	), u.controllers.CreateUserDetails)
	user.Put("/details", u.middlewares.HandleAuthWithRoles(
		utils.RoleUser,
	), u.controllers.UpdateUserDetails)
	user.Get("/details", u.middlewares.HandleAuthWithRoles(
		utils.RoleUser,
	), u.controllers.GetUser)
}

func NewUserRoute(
	controllers controllers.UserController,
	handler utils.RequestHandler,
	jwt middlewares.JwtMiddleware,
) UserRoute {
	return UserRoute{controllers, handler, jwt}
}
