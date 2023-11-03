package importar

import (
	"desafio1/service"
	"desafio1/tipos"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProdController interface {
	UpdateProd(ctx *gin.Context) error
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
