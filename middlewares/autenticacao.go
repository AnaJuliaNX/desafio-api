package middlewares

import (
	"desafio1/tipos"

	"github.com/gin-gonic/gin"
)

func BasicAutentication() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		tipos.Empresa.CNPJ: tipos.Empresa.Password,
	})
}
