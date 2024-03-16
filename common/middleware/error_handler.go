package middleware

import (
	"github.com/gin-gonic/gin"
	"go-demo/common"
	"net/http"
)

func NewErrorHandler(cfg *common.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		lastErr := c.Errors.Last()
		if lastErr == nil {
			return
		}

		var msg string
		if cfg.Debug {
			msg = lastErr.Error()
		} else {
			msg = "Server error"
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    lastErr.Type,
			"message": msg,
		})
	}
}
