package repositorio

import (
	"desafio1/tipos"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ProdutoRepositorio interface {
	SaveProd(tipos.Produto)
	UpdateProd(produto tipos.Produto)
	FindAllProd() []tipos.Produto
	FindProdId(produto tipos.Produto) tipos.Produto
	DeleteProd(produto tipos.Produto)
}

type bancoDedados struct {
	connection *gorm.DB
}

func NewProdutoRepositorio() ProdutoRepositorio {
	db, erro := gorm.Open("mysql", "testando:senhamudada@/testes?charset=utf8&parseTime=True&loc=Local")
	if erro != nil {
		log.Fatal("Falha ao estabeler a conexão com o banco de dados")
	}
	db.AutoMigrate(&tipos.Produto{}) //cria automaticamente os campos em uma tabela previamente criada
	return &bancoDedados{
		connection: db,
	}
}

func (db *bancoDedados) CloseDB() {
	erro := db.connection.Close()
	if erro != nil {
		log.Fatal("Erro ao encerrar a conexão com o banco de dados")
	}
}

func (db *bancoDedados) SaveProd(produto tipos.Produto) {
	db.connection.Create(&produto)
}

func (db *bancoDedados) UpdateProd(produto tipos.Produto) {
	db.connection.Save(&produto)
}

func (db *bancoDedados) FindAllProd() []tipos.Produto {
	var produtos []tipos.Produto
	db.connection.Find(&produtos)
	return produtos
}

func (db *bancoDedados) FindProdId(produto tipos.Produto) tipos.Produto {
	var encontrado tipos.Produto
	db.connection.First(&encontrado, produto.ID)
	return encontrado
}

func (db *bancoDedados) DeleteProd(produto tipos.Produto) {
	db.connection.Delete(&produto)
}
