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

func (c *controll) SaveEmp(ctx *gin.Context) error {
	var empresa tipos.Empresa
	erro := ctx.ShouldBindJSON(&empresa)
	if erro != nil {
		return erro
	}
	erro = validate.Struct(empresa)
	if erro != nil {
		return erro
	}
	c.service.SaveEmp(empresa)
	return nil
}

func (c *controll) UpdateEmp(ctx *gin.Context) error {
	var empresa tipos.Empresa
	erro := ctx.ShouldBind(&empresa)
	if erro != nil {
		return erro
	}
	id, erro := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if erro != nil {
		return erro
	}
	empresa.ID = id

	erro = validate.Struct(empresa)
	if erro != nil {
		return erro
	}
	c.service.UpdateEmp(empresa)
	return nil
}

func (c *controll) DeleteEmp(ctx *gin.Context) error {
	var empresa tipos.Empresa
	id, erro := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if erro != nil {
		return erro
	}
	empresa.ID = id
	c.service.DeleteEmp(empresa)
	return nil
}
