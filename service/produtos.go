package service

import (
	"desafio1/repositorio"
	"desafio1/tipos"
)

type ProdutoService interface {
	SaveProd(tipos.Produto) tipos.Produto
	UpdateProd(produto tipos.Produto)
	FindAllProd() []tipos.Produto
	FindProdId(produto tipos.Produto) (tipos.Produto, error)
	DeleteProd(produto tipos.Produto)
}

type produtoService struct {
	produtoRepositorio repositorio.ProdutoRepositorio
}

func New(repo repositorio.ProdutoRepositorio) ProdutoService {
	return &produtoService{
		produtoRepositorio: repo,
	}
}

func (service *produtoService) SaveProd(produto tipos.Produto) tipos.Produto {
	service.produtoRepositorio.SaveProd(produto)
	return produto
}

func (service *produtoService) UpdateProd(produto tipos.Produto) {
	service.produtoRepositorio.UpdateProd(produto)
}

func (service *produtoService) FindProdId(produto tipos.Produto) (tipos.Produto, error) {
	service.produtoRepositorio.FindProdId(produto)
	return produto, nil
}

func (service *produtoService) FindAllProd() []tipos.Produto {
	return service.produtoRepositorio.FindAllProd()
}

func (service *produtoService) DeleteProd(produto tipos.Produto) {
	service.produtoRepositorio.DeleteProd(produto)
}
