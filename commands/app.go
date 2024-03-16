package commands

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"go-demo/common"
	"go.uber.org/config"
	"os"
)

func setConfigArg(context *cli.Context, cfg *common.Config, logger *log.Logger) {

	file := context.String("config")
	if file == "" {
		defaultFile := "config/config.yaml"
		if _, err := os.Stat(defaultFile); err != nil {
			return
		}
		file = defaultFile
	} else if _, err := os.Stat(file); err != nil {
		panic(err.Error())
	}
	logger.Debugf("Read config file: %s", file)
	provider, err := config.NewYAML(config.File(file))
	provider.Get("").Populate(cfg)
	if err != nil {
		panic(err) // handle error
	}
}

func setLogLevelArg(context *cli.Context, logger *log.Logger) {
	level := context.String("level")
	if level == "" {
		return
	}
	logLevel, err := log.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	logger.SetLevel(logLevel)
}

func NewCommandApplication(commands []*cli.Command, cfg *common.Config, logger *log.Logger) *cli.App {
	return &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
			},
			&cli.StringFlag{
				Name:    "level",
				Aliases: []string{"l"},
				Usage:   "Set log level, default `info`.",
			},
		},
		Before: func(context *cli.Context) error {
			setLogLevelArg(context, logger)
			cwd, _ := os.Getwd()
			logger.Debugf("Workdir: %s", cwd)
			setConfigArg(context, cfg, logger)
			return nil
		},
		Commands: commands,
	}
}
