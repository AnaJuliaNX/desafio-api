package service

type LoginService interface {
	Login(cnpj int64, password string) bool
}

type loginService struct {
	authorizedCnpj     int64
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedCnpj:     123,
		authorizedPassword: "teste",
	}
}

func (service *loginService) Login(cnpj int64, password string) bool {
	return service.authorizedCnpj == cnpj &&
		service.authorizedPassword == password
}
