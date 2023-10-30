package service

import "desafio1/tipos"

type LoginService interface {
	Login(cnpj string, password string) bool
}

type loginService struct {
	authorizedCnpj     string
	authorizedPassword string
}

// Consulta o banco de dados para pegar os dados previamente cadastrados
func NewLoginService() LoginService {

	var empresa tipos.Empresa
	return &loginService{
		authorizedCnpj:     empresa.CNPJ,
		authorizedPassword: empresa.Password,
	}
}

func (service *loginService) Login(cnpj string, password string) bool {
	return service.authorizedCnpj == cnpj &&
		service.authorizedPassword == password
}
