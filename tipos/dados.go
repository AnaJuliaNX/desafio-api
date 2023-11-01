package tipos

// Será usado para gerar o token e fazer o login
type Empresa struct {
	ID       int64  `json:"id"`
	CNPJ     string `json:"cnpj"`
	Password string `json:"password"`
}

// Cadastrar o produto
type Produto struct {
	ID            int64  `json:"cProd"`
	Codigo        string `json:"cEAN"`
	Descricao     string `json:"xProd"`
	UniMedida     string `json:"uCom"`
	Quantidade    string `json:"qCom"`
	ValorUnitario string `json:"vUnCom"`
}

// Cadastrar para conferir com os dados da XML
type Destinatario struct {
	ID         int64  `json:"id"`
	CNPJ       string `json:"cnpj"`
	Logradouro string `json:"lgr"`
	Numero     string `json:"nro"`
	Bairro     string `json:"bairro"`
	Municipio  string `json:"xMun"`
	UF         string `json:"uf"`
	CEP        string `json:"cep"`
	Pais       string `json:"xPais"`
	Telefone   string `json:"fone"`
}

type Emitente struct {
	CNPJ string
}
