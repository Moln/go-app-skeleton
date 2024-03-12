package commands

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"go-demo/common"
)

func NewServeCommand(cfg *common.Config, http *gin.Engine, logger *log.Logger) *cli.Command {
	return &cli.Command{
		Name:  "serve",
		Usage: "Run http server",
		Action: func(cCtx *cli.Context) error {
			logger.Infof("Running http server %s", cfg.Http.Addr)
			http.Run(cfg.Http.Addr)
			return nil
		},
	}
}
