package main

import (
	"desafio1/controller"
	"desafio1/repositorio"
	"desafio1/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	prodRepositorio repositorio.ProdutoRepositorio = repositorio.NewProdutoRepositorio()
	destRepositorio repositorio.DestRepositorio    = repositorio.NewDestRepositorio()
	prodService     service.ProdutoService         = service.New(prodRepositorio)
	destService     service.DestService            = service.News(destRepositorio)
	loginService    service.LoginService           = service.NewLoginService()
	jwtService      service.JWTService             = service.NewJWTService()

	prodController  controller.ProdController  = controller.New(prodService)
	destController  controller.DestController  = controller.News(destService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func main() {
	//para testar se ta funcionando
	server := gin.Default()
	server.GET("/teste", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!",
		})
	})

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	server.Run(":" + port)
}
