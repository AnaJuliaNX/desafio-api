package controller

import (
	"desafio1/service"
	"desafio1/tipos"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func NewLoginController(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credenciais tipos.Empresa
	erro := ctx.ShouldBind(&credenciais)
	if erro != nil {
		return ""
	}

	autenticado := controller.loginService.Login(credenciais.CNPJ, credenciais.Password)
	if autenticado {
		return controller.jwtService.GenerateToken(credenciais.CNPJ, credenciais.Password, true)
	}
	return ""
}
