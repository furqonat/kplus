package controllers

import (
	"go.uber.org/fx"
	v1 "kplus.com/controllers/v1"
)

var Module = fx.Options(
	fx.Provide(v1.NewAuthController),
)
