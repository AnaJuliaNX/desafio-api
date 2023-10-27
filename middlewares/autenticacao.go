package middlewares

import (
	"desafio1/tipos"

	"github.com/gin-gonic/gin"
)

// Para fazer a autenticação dos dados inseridos, user(cnpj), senha
func BasicAutentication() gin.HandlerFunc {
	var empresa tipos.Empresa
	return gin.BasicAuth(gin.Accounts{
		empresa.CNPJ: empresa.Password,
	})
}
