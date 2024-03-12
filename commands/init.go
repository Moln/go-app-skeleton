package commands

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"go-demo/common"
	"go.uber.org/config"
	"os"
)

func NewCommandApplication(commands []*cli.Command, cfg *common.Config, logger *log.Logger) *cli.App {
	return &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
			},
		},
		Before: func(context *cli.Context) error {
			file := context.String("config")
			if file == "" {
				defaultFile := "config/config.yaml"
				if _, err := os.Stat(defaultFile); err == nil {
					file = defaultFile
				}
				return nil
			} else if _, err := os.Stat(file); err != nil {
				panic(err.Error())
			}
			logger.Infof("Read config file: %s", file)
			provider, err := config.NewYAML(config.File(file))
			provider.Get("").Populate(cfg)
			if err != nil {
				panic(err) // handle error
			}

			return nil
		},
		Commands: commands,
	}
}
