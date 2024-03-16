package di

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	log "github.com/sirupsen/logrus"
	"go-demo/commands"
	"go-demo/common"
	"go-demo/common/middleware"
	"go-demo/common/router"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Module = fx.Module("di",
	fx.Provide(
		common.NewConfig,
		NewGorm,
		NewLogger,
		fx.Annotate(
			NewHttpServer,
			fx.ParamTags(`group:"routes"`, `group:"middlewares"`),
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
		fx.Annotate(
			commands.NewConfigCommand,
			fx.ResultTags(`group:"commands"`),
		),
		fx.Annotate(
			middleware.NewSessionMiddleware,
			fx.ResultTags(`group:"middlewares"`),
		),
		fx.Annotate(
			middleware.NewErrorHandler,
			fx.ResultTags(`group:"middlewares"`),
		),
	),
)

func NewLogger() *log.Logger {
	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{})
	logger.SetLevel(log.DebugLevel)
	//logger.AddHook(&GlobalHook{})
	return logger
}

func NewGorm(conf *common.Config) *gorm.DB {

	var driver gorm.Dialector
	dsn := conf.Db.Dsn
	switch conf.Db.Driver {
	case "sqlite":
		driver = sqlite.Open(dsn)
	case "mysql":
		driver = mysql.Open(dsn)
	default:
		panic(fmt.Sprintf(`Invalid driver name: '%s'`, conf.Db.Driver))
	}
	db, err := gorm.Open(driver, &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + dsn + "\n" + err.Error())
	}

	return db
}

func NewHttpServer(routes []router.Route, middlewares gin.HandlersChain) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	for _, middleware := range middlewares {
		r.Use(middleware)
	}
	for _, route := range routes {
		route.Register(r)
	}

	return r
}
