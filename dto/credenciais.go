package dto

type Credenciais struct {
	Cnpj     string `form:"cnpj"`
	Password string `form:"password"`
}
