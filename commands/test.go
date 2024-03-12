package commands

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"go-demo/common"
)

func NewTestCommand(cfg *common.Config, http *gin.Engine, logger *log.Logger) *cli.Command {
	return &cli.Command{
		Name:  "test",
		Usage: "test...",
		Action: func(cCtx *cli.Context) error {
			logger.Infof("Test %v", cfg.Http.Addr)
			return nil
		},
	}
}
