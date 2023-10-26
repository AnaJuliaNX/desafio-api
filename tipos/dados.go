package tipos

// Será usado para gerar o token e fazer o login
type Empresa struct {
	ID       int64  `json:"id"`
	CNPJ     int64  `json:"cnpj"`
	Password string `json:"password"`
}

// Cadastrar o produto
type Produto struct {
	ID            int64   `json:"cProd"`
	Código        int64   `json:"cEAN"`
	Descricao     string  `json:"xProd"`
	UniMedida     string  `json:"uCom"`
	Quantidade    int64   `json:"qCom"`
	ValorUnitario float64 `json:"vUnCom"`
}

// Cadastrar para conferir com os dados da XML
type Destinatario struct {
	ID         int64  `json:"id"`
	Logradouro string `json:"lgr"`
	Numero     int64  `json:"nro"`
	Bairro     string `json:"bairro"`
	Municipio  string `json:"xMun"`
	UF         string `json:"uf"`
	CEP        int64  `json:"cep"`
	Pais       string `json:"xPais"`
	Telefone   int64  `json:"fone"`
}
