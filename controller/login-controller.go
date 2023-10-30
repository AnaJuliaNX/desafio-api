package controller

import (
	"desafio1/dto"
	"desafio1/service"

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

// Função para fazer o login com as credenciais de acesso e gerar um token
func (controller *loginController) Login(ctx *gin.Context) string {
	var credenciais dto.Credenciais
	erro := ctx.ShouldBind(&credenciais)
	if erro != nil {
		return ""
	}
	autenticado := controller.loginService.Login(credenciais.Cnpj, credenciais.Password)
	if autenticado {
		return controller.jwtService.GenerateToken(credenciais.Cnpj, credenciais.Password, true)
	}
	return ""
}
