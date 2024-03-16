package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"go-demo/di"
	"go-demo/module/app"
	"go.uber.org/dig"
	"go.uber.org/fx"
	"os"
	"reflect"
)

func main() {
	//cmd.Execute()
	app := fx.New(
		di.Module,
		app.Module,
		//fx.Invoke(func(app *cli.App, logger *log.Logger, lifecycle fx.Lifecycle) {
		//	log.Info("22")
		//	if err := app.Run(os.Args); err != nil {
		//		log.Fatal(err)
		//	}
		//}),
	)
	rApp := reflect.ValueOf(*app)
	fieldContainer := rApp.FieldByName("container")
	container := (*dig.Container)(fieldContainer.UnsafePointer())
	container.Provide(func() *dig.Container {
		return container
	})
	err := container.Invoke(func(app *cli.App, logger *log.Logger, cc *dig.Container) {
		if err := app.Run(os.Args); err != nil {
			logger.Fatal(err)
		}
	})
	if err != nil {
		panic(err)
	}
	//containerValue := v.FieldByName("container")
	//container := containerValue.Interface()
}

type Test struct {
	A int
	b int
}
