package routes

import "go.uber.org/fx"

var Module = fx.Options()

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes() Routes {
	return Routes{}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
