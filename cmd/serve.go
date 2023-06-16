package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-demo/di"
	"go-demo/module/app"
	"go.uber.org/fx"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		fx.New(
			fx.Provide(
				fx.Annotated{
					Name:   "config",
					Target: func() string { return cfgFile },
				},
			),
			di.Module,
			app.Module,
			fx.Invoke(func(r *gin.Engine) {
				r.Run(":18080")
			}),
		)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
