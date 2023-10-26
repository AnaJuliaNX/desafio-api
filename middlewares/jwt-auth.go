package middlewares

import (
	"desafio1/service"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Autorizacao() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BERAER_SCHEMA = "Beraer"
		authHeader := c.GetHeader("Autorização")
		tokenString := authHeader[len(BERAER_SCHEMA):]
		token, erro := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Cnpj]: ", claims["cnpj"])
			log.Println("Claims[Senha]: ", claims["senha"])
			log.Println("Claims[admin]: ", claims["admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpireAt]: ", claims["exp"])
		} else {
			log.Println(erro)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
