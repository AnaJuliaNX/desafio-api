package middlewares

import (
	"github.com/gin-gonic/gin"
)

func BasicAutentication() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"teste": "teste",
	})
}
