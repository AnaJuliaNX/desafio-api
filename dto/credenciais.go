package dto

// credenciais de acesso
type Credenciais struct {
	Cnpj     string `form:"cnpj"`
	Password string `form:"password"`
}
