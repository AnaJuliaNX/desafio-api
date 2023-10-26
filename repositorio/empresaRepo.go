package repositorio

import (
	"desafio1/tipos"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type EmpRepositorio interface {
	SaveEmp(tipos.Empresa)
	UpdateEmp(emp tipos.Empresa)
	DeleteEmp(emp tipos.Empresa)
}

type banCodeDados struct {
	connection *gorm.DB
}

func NewEmpRepositorio() EmpRepositorio {
	db, erro := gorm.Open("mysql", "testando:senhamudada@/testes?charset=utf8&parseTime=True&loc=Local")
	if erro != nil {
		log.Fatal("Falha ao estabeler a conexão com o banco de dados")
	}
	db.AutoMigrate(&tipos.Empresa{})
	return &banCodeDados{
		connection: db,
	}
}

func (db *bancodeDados) CloseDb() {
	erro := db.connection.Close()
	if erro != nil {
		log.Fatal("Erro ao encerrar a conexão com o banco de dados")
	}
}

func (db *banCodeDados) SaveEmp(emp tipos.Empresa) {
	db.connection.Create(&emp)
}

func (db *banCodeDados) UpdateEmp(emp tipos.Empresa) {
	db.connection.Save(&emp)
}

func (db *banCodeDados) DeleteEmp(emp tipos.Empresa) {
	db.connection.Delete(&emp)
}
