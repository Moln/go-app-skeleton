package app

import (
	"go-demo/common/router"
	"go-demo/module/app/controller"
	"go.uber.org/fx"
)

var Module = fx.Module("app",
	fx.Provide(
		controller.NewAppController,
		controller.NewAuthController,
		router.AsRoute(NewAppRoute),
	),
)
