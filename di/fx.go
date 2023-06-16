package di

import (
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	log "github.com/sirupsen/logrus"
	"go-demo/module/route"
	"go.uber.org/config"
	"go.uber.org/fx"
	gorm "gorm.io/gorm"
	"os"
)

var Module = fx.Module("di",
	fx.Provide(
		NewGorm,
		NewLogger,
		fx.Annotate(
			NewHttpServer,
			fx.ParamTags(`group:"routes"`),
		),
		fx.Annotate(
			NewConfigProvider,
			fx.ParamTags(`name:"config"`),
		),
	),
)


func NewLogger() *log.Logger {
	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{})
	//logger.AddHook(&GlobalHook{})
	return logger
}

func NewGorm(conf *config.YAML) *gorm.DB {

	dsn := conf.Get("db").String()
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

func NewConfigProvider(file string) *config.YAML {
	_, err := os.Stat(file)
	if err != nil {
		panic(err) // handle error
	}
	provider, err := config.NewYAML(config.File(file))
	if err != nil {
		panic(err) // handle error
	}

	return provider
}
