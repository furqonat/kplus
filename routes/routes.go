package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewAuthRoute),
	fx.Provide(NewUserRoute),
	fx.Provide(NewTransactionRoute),
	fx.Provide(NewInstallmentRoute),
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	authRoute AuthRoute,
	userRoute UserRoute,
	transactionRoute TransactionRoute,
	installmentRoute InstallmentRoute,
) Routes {
	return Routes{
		authRoute,
		userRoute,
		transactionRoute,
		installmentRoute,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
