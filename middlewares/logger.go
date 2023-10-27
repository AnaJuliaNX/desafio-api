package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(parametro gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s \n",
			parametro.ClientIP,   //endereço de IP do cliente
			parametro.TimeStamp,  //data e hora em que a solicitação foi feita
			parametro.Method,     //método HTTP da solicitação (Post, Get, Put, Delete)
			parametro.Path,       //caminho (URL) da solicitação
			parametro.StatusCode, //código de status HTTP da resposta
			parametro.Latency,    //duração da solicitação
		)
	})
}
