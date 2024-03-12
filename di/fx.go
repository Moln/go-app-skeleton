package di

import (
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	log "github.com/sirupsen/logrus"
	"go-demo/commands"
	"go-demo/common"
	"go-demo/module/route"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Module("di",
	fx.Provide(
		common.NewConfig,
		NewGorm,
		NewLogger,
		fx.Annotate(
			NewHttpServer,
			fx.ParamTags(`group:"routes"`),
		),
		fx.Annotate(
			commands.NewCommandApplication,
			fx.ParamTags(`group:"commands"`),
		),

		fx.Annotate(
			commands.NewServeCommand,
			fx.ResultTags(`group:"commands"`),
		),
		fx.Annotate(
			commands.NewTestCommand,
			fx.ResultTags(`group:"commands"`),
		),
	),
)

func NewLogger() *log.Logger {
	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{})
	//logger.AddHook(&GlobalHook{})
	return logger
}

func NewGorm(conf *common.Config) *gorm.DB {

	dsn := conf.Db.Dsn
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + dsn + "\n" + err.Error())
	}

	return db
}

func NewHttpServer(routes []route.Route) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	for _, route := range routes {
		route.Register(r)
	}

	return r
}
