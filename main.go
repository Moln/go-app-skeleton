package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"go-demo/di"
	"go-demo/module/app"
	"go.uber.org/fx"
	"os"
)

func main() {
	//cmd.Execute()
	fx.New(
		di.Module,
		app.Module,
		fx.Invoke(func(app *cli.App, logger *log.Logger) {
			if err := app.Run(os.Args); err != nil {
				log.Fatal(err)
			}
		}),
	)
}
