package repositorio

import (
	"desafio1/tipos"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DestRepositorio interface {
	SaveDest(tipos.Destinatario)
	UpdateDest(dest tipos.Destinatario)
	FindAllDest() []tipos.Destinatario
	DeleteDest(dest tipos.Destinatario)
}

// Para fazer a conexão com o banco de dados, o escolhido foi o mysql
type bancodeDados struct {
	connection *gorm.DB
}

func NewDestRepositorio() DestRepositorio {
	db, erro := gorm.Open("mysql", "testando:senhamudada@/testes?charset=utf8&parseTime=True&loc=Local")
	if erro != nil {
		log.Fatal("Falha ao estabeler a conexão com o banco de dados")
	}
	db.AutoMigrate(&tipos.Destinatario{}) //cria automaticamente os campos em uma tabela previamente criada
	return &bancodeDados{
		connection: db,
	}
}

func (db *bancodeDados) Closedb() {
	erro := db.connection.Close()
	if erro != nil {
		log.Fatal("Erro ao encerrar a conexão com o banco de dados")
	}
}

func (db *bancodeDados) SaveDest(dest tipos.Destinatario) {
	db.connection.Create(&dest)
}

func (db *bancodeDados) UpdateDest(dest tipos.Destinatario) {
	db.connection.Save(&dest)
}

func (db *bancodeDados) FindAllDest() []tipos.Destinatario {
	var dests []tipos.Destinatario
	db.connection.Find(&dests)
	return dests
}

func (db *bancodeDados) DeleteDest(dest tipos.Destinatario) {
	db.connection.Delete(&dest)
}
