package commands

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"go-demo/common"
	"go.uber.org/dig"
)

func NewServeCommand(container *dig.Container) *cli.Command {
	return &cli.Command{
		Name:  "serve",
		Usage: "Run http server",
		Action: func(cCtx *cli.Context) error {
			container.Invoke(func(http *gin.Engine, cfg *common.Config, logger *log.Logger) {
				logger.Infof("Running http server %s", cfg.Http.Addr)
				http.Run(cfg.Http.Addr)
			})
			return nil
		},
	}
}
