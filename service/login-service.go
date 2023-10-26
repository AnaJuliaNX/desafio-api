package service

import "desafio1/tipos"

type LoginService interface {
	Login(cnpj int64, password string) bool
}

type loginService struct {
	authorizedCnpj     int64
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedCnpj:     tipos.Empresa.CNPJ,
		authorizedPassword: tipos.Empresa.Password,
	}
}

func (service *loginService) Login(cnpj int64, password string) bool {
	return service.authorizedCnpj == cnpj &&
		service.authorizedPassword == password
}
