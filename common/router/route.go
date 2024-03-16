package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

type Route interface {
	Register(r *gin.Engine)
}
