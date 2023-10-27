package controller

import (
	"desafio1/service"
	"desafio1/tipos"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProdController interface {
	SaveProd(cxt *gin.Context) error
	UpdateProd(ctx *gin.Context) error
	FindAllProd() []tipos.Produto
	DeleteProd(ctx *gin.Context) error
}

type controller struct {
	service service.ProdutoService
}

var validate *validator.Validate

func New(service service.ProdutoService) ProdController {
	validate = validator.New()
	return &controller{
		service: service,
	}
}

// Função para cadastrar um novo produto
func (c *controller) SaveProd(ctx *gin.Context) error {
	var produto tipos.Produto
	erro := ctx.ShouldBindJSON(&produto)
	if erro != nil {
		return erro
	}
	erro = validate.Struct(produto)
	if erro != nil {
		return erro
	}
	c.service.SaveProd(produto)
	return nil
}

// Função para atualizar um produto já cadastrado, busca pelo ID na rota
func (c *controller) UpdateProd(ctx *gin.Context) error {
	var produto tipos.Produto
	erro := ctx.ShouldBind(&produto)
	if erro != nil {
		return erro
	}
	id, erro := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if erro != nil {
		return erro
	}
	produto.ID = id

	erro = validate.Struct(produto)
	if erro != nil {
		return erro
	}
	c.service.UpdateProd(produto)
	return nil
}

// Função para listar todos os produtos cadastrados
func (c *controller) FindAllProd() []tipos.Produto {
	return c.service.FindAllProd()
}

// Função para deletar um produto, busca pelo Id na rota
func (c *controller) DeleteProd(ctx *gin.Context) error {
	var produto tipos.Produto
	id, erro := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if erro != nil {
		return erro
	}
	produto.ID = id
	c.service.DeleteProd(produto)
	return nil
}
