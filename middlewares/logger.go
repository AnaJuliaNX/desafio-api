package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(parametro gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s \n",
			parametro.ClientIP,
			parametro.TimeStamp,
			parametro.Method,
			parametro.Path,
			parametro.StatusCode,
			parametro.Latency,
		)
	})
}
