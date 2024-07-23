package routes

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"kplus.com/controllers"
	"kplus.com/middlewares"
	"kplus.com/utils"
)

type AuthRoute struct {
	controllers controllers.AuthController
	handler     utils.RequestHandler
	middlewares middlewares.JwtMiddleware
}

func (a AuthRoute) Setup() {

	log.Println("setting up auth routes")
	auth := a.handler.App.Group("/auth").Use(limiter.New(
		limiter.Config{
			Max:        100,
			Expiration: 60 * time.Second,
		}),
		cors.New(cors.Config{
			AllowOrigins:     "http://localhost:8080, http://localhost:3000",
			AllowMethods:     "*",
			AllowCredentials: true,
		}),
	)
	auth.Post("/signIn", a.controllers.SignIn)
	auth.Post("/signUp", a.controllers.SignUp)
	auth.Get("/refresh", a.middlewares.HandleAuthWithRoles(
		utils.RoleUser,
	), a.controllers.RefreshToken)
}

func NewAuthRoute(
	controllers controllers.AuthController,
	handler utils.RequestHandler,
	jwt middlewares.JwtMiddleware,
) AuthRoute {
	return AuthRoute{controllers, handler, jwt}
}
