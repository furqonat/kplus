package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewAuthRoute),
	fx.Provide(NewUserRoute),
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	authRoute AuthRoute,
	userRoute UserRoute,
) Routes {
	return Routes{
		authRoute,
		userRoute,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
