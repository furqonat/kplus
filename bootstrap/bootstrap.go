package bootstrap

import (
	"context"

	"go.uber.org/fx"
	"kplus.com/controllers"
	"kplus.com/middlewares"
	"kplus.com/routes"
	"kplus.com/services"
	"kplus.com/utils"
)

var Module = fx.Options(
	controllers.Module,
	utils.Module,
	middlewares.Module,
	routes.Module,
	services.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler utils.RequestHandler,
	routes routes.Routes,
	logger utils.Logger,
	env utils.Env,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				routes.Setup()
				port := env.ServerPort
				if err := handler.App.Listen(":" + port); err != nil {
					logger.Panicf("failed to start server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("shutting down server")
			return nil
		},
	})
}
