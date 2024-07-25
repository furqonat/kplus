package services

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAuthService),
	fx.Provide(NewUserService),
	fx.Provide(NewTransactionService),
	fx.Provide(NewInstallmentService),
)
