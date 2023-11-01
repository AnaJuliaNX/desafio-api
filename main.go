package main

import (
	"desafio1/controller"
	"desafio1/middlewares"
	"desafio1/repositorio"
	"desafio1/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	empRepositorio  repositorio.EmpRepositorio     = repositorio.NewEmpRepositorio()
	prodRepositorio repositorio.ProdutoRepositorio = repositorio.NewProdutoRepositorio()
	destRepositorio repositorio.DestRepositorio    = repositorio.NewDestRepositorio()

	empService   service.EmpService     = service.NeW(empRepositorio)
	prodService  service.ProdutoService = service.New(prodRepositorio)
	destService  service.DestService    = service.News(destRepositorio)
	loginService service.LoginService   = service.NewLoginService()
	jwtService   service.JWTService     = service.NewJWTService()

	empController   controller.EmpController   = controller.NewS(empService)
	destController  controller.DestController  = controller.News(destService)
	prodController  controller.ProdController  = controller.New(prodService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func main() {
	//para testar se ta funcionando
	server := gin.Default()
	// server.GET("/teste", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "OK!",
	// 	})
	// })

	//Endpoint de login: autenticação + token criado
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

	apiRouters := server.Group("/api", middlewares.Autorizacao())
	{
		//ROTAS PARA A EMPRESA
		apiRouters.POST("/empresa", func(ctx *gin.Context) {
			erro := empController.SaveEmp(ctx)
			if erro != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Empresa adicionada com sucesso"})
			}
		})
		apiRouters.PUT("/empresa/:id", func(ctx *gin.Context) {
			erro := empController.UpdateEmp(ctx)
			if erro != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Empresa atualizada com sucesso"})
			}
		})
		apiRouters.DELETE("/empresa", func(ctx *gin.Context) {
			erro := empController.DeleteEmp(ctx)
			if erro != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Empresa deletada com sucesso"})
			}
		})

		//ROTAS PARA O DESTINATARIO
		apiRouters.GET("/destinatarios", func(ctx *gin.Context) {
			ctx.JSON(200, destController.FindAllDest())
		})
		apiRouters.POST("/destinatarios", func(ctx *gin.Context) {
			erro := destController.SaveDest(ctx)
			if erro != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Destinatario adicionado com sucesso!"})
			}
		})
		apiRouters.PUT("/destinatarios/:id", func(ctx *gin.Context) {
			erro := destController.UpdateDest(ctx)
			if erro != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Destinatario atualizado com sucesso"})
			}
		})
		apiRouters.DELETE("/destinatarios/:id", func(ctx *gin.Context) {
			erro := destController.DeleteDest(ctx)
			if erro != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Destinatario deletado com sucesso"})
			}
		})

		//ROTAS PARA O PRODUTO
		apiRouters.GET("/produtos", func(ctx *gin.Context) {
			ctx.JSON(200, prodController.FindAllProd())
		})
		apiRouters.POST("/produtos", func(ctx *gin.Context) {
			erro := prodController.SaveProd(ctx)
			if erro != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Produto adicionado com sucesso"})
			}
		})
		apiRouters.PUT("/produtos/:id", func(ctx *gin.Context) {
			erro := prodController.UpdateProd(ctx)
			if erro != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Produto atualizado com sucesso"})
			}
		})
		apiRouters.DELETE("/produtos/:id", func(ctx *gin.Context) {
			erro := prodController.DeleteProd(ctx)
			if erro != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Produto deletado com sucesso"})
			}
		})
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
}
