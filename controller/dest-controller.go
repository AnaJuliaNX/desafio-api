package controller

import (
	"desafio1/service"
	"desafio1/tipos"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DestController interface {
	SaveDest(cxt *gin.Context) error
	UpdateDest(ctx *gin.Context) error
	FindAllDest() []tipos.Destinatario
	DeleteDest(ctx *gin.Context) error
}

type controler struct {
	service service.DestService
}

var validat *validator.Validate

func News(service service.DestService) DestController {
	validate = validator.New()
	return &controler{
		service: service,
	}
}

// Função para cadastrar um destinatario
func (c *controler) SaveDest(ctx *gin.Context) error {
	var destinatario tipos.Destinatario
	erro := ctx.ShouldBindJSON(&destinatario)
	if erro != nil {
		return erro
	}
	erro = validate.Struct(destinatario)
	if erro != nil {
		return erro
	}
	c.service.SaveDest(destinatario) //função que vai salvar o detinatário
	return nil
}

// Função para atualizar oui editar o destinatário, busca pelo ID na rota
func (c *controler) UpdateDest(ctx *gin.Context) error {
	var destinatario tipos.Destinatario
	erro := ctx.ShouldBind(&destinatario)
	if erro != nil {
		return erro
	}
	id, erro := strconv.ParseInt(ctx.Param("id"), 0, 0) //pega o ID da rota
	if erro != nil {
		return erro
	}
	destinatario.ID = id

	erro = validate.Struct(destinatario) //faz a validação dos dados
	if erro != nil {
		return erro
	}
	c.service.UpdateDest(destinatario) //se deu tudo certo na validação atualiza os dados do destinatario
	return nil
}

// Função para achar todos os destinatarios cadastrados
func (c *controler) FindAllDest() []tipos.Destinatario {
	return c.service.FindAllDest()
}

// Função para deletar um destinatario, busca pelo ID na rota
func (c *controler) DeleteDest(ctx *gin.Context) error {
	var destinatario tipos.Destinatario
	id, erro := strconv.ParseInt(ctx.Param("id"), 0, 0) //pega o id na rota
	if erro != nil {
		return erro
	}
	destinatario.ID = id
	c.service.DeleteDest(destinatario) //deleta os dados de um destinatario especificado pelo ID
	return nil
}
