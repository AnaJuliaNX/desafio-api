package service

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type LoginService interface {
	Login(cnpj string, password string) bool
}

type loginService struct {
	db *sql.DB
}

// Consulta o banco de dados para pegar os dados previamente cadastrados
func NewLoginService() LoginService {
	db, erro := sql.Open("mysql", "testando:senhamudada@/testes?charset=utf8&parseTime=True&loc=Local")
	if erro != nil {
		log.Fatalf("Erro ao estabelecer a conexão com o banco de dados")
		return nil
	}
	return &loginService{db}
}

func (service *loginService) Login(cnpj string, password string) bool {
	lines, erro := service.db.Query("select cnpj, password from empresas where cnpj = ? and password = ?", cnpj, password)
	if erro != nil {
		fmt.Println(erro)
		log.Fatalf("Erro ao ")
	}
	var authorizedCnpj, authorizedPassword string
	for lines.Next() {
		erros := lines.Scan(&authorizedCnpj, &authorizedPassword)
		if erros != nil {
			log.Fatal("Erro ao realizar a consulta no banco de dados")
			return false
		}
	}
	return cnpj == authorizedCnpj && password == authorizedPassword
}
