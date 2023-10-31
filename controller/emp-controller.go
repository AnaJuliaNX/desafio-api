package controller

import (
	"desafio1/service"
	"desafio1/tipos"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EmpController interface {
	SaveEmp(cxt *gin.Context) error
	UpdateEmp(ctx *gin.Context) error
	DeleteEmp(ctx *gin.Context) error
}

type controll struct {
	service service.EmpService
}

var valida *validator.Validate

func NewS(service service.EmpService) EmpController {
	validate = validator.New()
	return &controll{
		service: service,
	}
}

// Função para cadastrar uma empresa
func (c *controll) SaveEmp(ctx *gin.Context) error {
	var empresa tipos.Empresa
	erro := ctx.ShouldBindJSON(&empresa)
	if erro != nil {
		return erro
	}
	erro = validate.Struct(empresa) //faz a validação dos dados
	if erro != nil {
		return erro
	}
	c.service.SaveEmp(empresa) //se deu tudo certo na validação salva os dados da empresa
	return nil
}

// Função para atualizar uma empresa ou alterar, busca pelo ID na rota
func (c *controll) UpdateEmp(ctx *gin.Context) error {
	var empresa tipos.Empresa
	erro := ctx.ShouldBind(&empresa)
	if erro != nil {
		return erro
	}
	id, erro := strconv.ParseInt(ctx.Param("id"), 0, 0) //Busco a empresa especificada pelo ID na rota
	if erro != nil {
		return erro
	}
	empresa.ID = id

	erro = validate.Struct(empresa) //faz a validação dos dados
	if erro != nil {
		return erro
	}
	c.service.UpdateEmp(empresa) //se deu tudo certo na validação salva os dados atualizados
	return nil
}

// Função para deletar uma empresa, busca pelo ID na rota
func (c *controll) DeleteEmp(ctx *gin.Context) error {
	var empresa tipos.Empresa
	id, erro := strconv.ParseInt(ctx.Param("id"), 0, 0) //Busca a empresa especificada pelo ID na rota
	if erro != nil {
		return erro
	}
	empresa.ID = id
	c.service.DeleteEmp(empresa) //deleta a empresa especificada
	return nil
}
