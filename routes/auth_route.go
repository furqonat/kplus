package routes

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	v1 "kplus.com/controllers/v1"
	"kplus.com/utils"
)

type AuthRoute struct {
	logger      utils.Logger
	controllers v1.AuthController
	handler     utils.RequestHandler
}

func (a AuthRoute) Setup() {

	a.logger.Info("setting up auth routes")
	auth := a.handler.App.Group("/auth").Use(limiter.New(
		limiter.Config{
			Max:        100,
			Expiration: 60 * time.Second,
		}),
		cors.New(cors.Config{
			AllowOrigins:     "*",
			AllowMethods:     "*",
			AllowCredentials: true,
		}),
	)
	auth.Post("/signIn", a.controllers.SignIn)
	auth.Post("/signUp", a.controllers.SignUp)
	auth.Post("/refreshToken", a.controllers.RefreshToken)
}
