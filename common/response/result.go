package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResult struct {
	Message string
}

func Error(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusForbidden, ErrorResult{
		Message: message,
	})
}
