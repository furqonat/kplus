package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"kplus.com/bootstrap"
	"kplus.com/utils"
)

func main() {
	godotenv.Load()
	logger := utils.GetLogger().GetFxLogger()
	fx.New(bootstrap.Module, fx.Logger(logger)).Run()
}
