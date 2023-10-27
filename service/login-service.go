package service

type LoginService interface {
	Login(cnpj string, password string) bool
}

type loginService struct {
	authorizedCnpj     string
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedCnpj:     "123",
		authorizedPassword: "teste",
	}
}

func (service *loginService) Login(cnpj string, password string) bool {
	return service.authorizedCnpj == cnpj &&
		service.authorizedPassword == password
}
