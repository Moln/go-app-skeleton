package app

import (
	"go-demo/module/app/controller"
	"go-demo/module/app/router"
	"go-demo/module/route"
	"go.uber.org/fx"
)

var Module = fx.Module("app",
	fx.Provide(
		controller.NewAppController,
		controller.NewAuthController,
		route.AsRoute(router.NewAppRoute),
	),
)
