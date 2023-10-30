package middlewares

import (
	"desafio1/service"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Função para gerar o token e fazer a validação dele
func Autorizacao() gin.HandlerFunc {
	return func(c *gin.Context) {
		//extrai o token  da solicitação http
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		//ffaz a validação do token
		token, erro := service.NewJWTService().ValidateToken(tokenString)

		//Se o token for válido extra essas informações dele
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Cnpj]: ", claims["cnpj"])
			log.Println("Claims[Password]: ", claims["password"])
			log.Println("Claims[admin]: ", claims["admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpireAt]: ", claims["exp"])
		} else {
			//Se for invalido exibe o erro, a mensagem e interrompe a execução
			log.Println(erro)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
