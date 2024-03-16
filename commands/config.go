package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go-demo/common"
)

func NewConfigCommand(cfg *common.Config) *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "Show config list.",
		Action: func(c *cli.Context) error {
			fmt.Printf("Config: %+v", cfg)
			return nil
		},
	}
}
