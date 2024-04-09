package middleware

import (
	"github.com/gin-gonic/gin"
	"go-demo/common"
	"net/http"
	"runtime/debug"
)

func NewErrorHandler(cfg *common.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
				if str, ok := err.(string); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code":    http.StatusInternalServerError,
						"message": str,
					})
					return
				}
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "Server error",
				})
			}
		}()

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
