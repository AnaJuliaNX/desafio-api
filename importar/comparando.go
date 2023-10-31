package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Allmap struct {
	Emits []Emit `xml:"emit"`
	Dests []Dest `xml:"dest"`
	Prods []Prod `xml:"prod"`
}

type Emit struct {
	CNPJ       int64  `xml:"cnpj"`
	Nome       string `xml:"xNome"`
	Logradouro string `xml:"lgr"`
	Numero     int64  `xml:"nro"`
	Bairro     string `xml:"bairro"`
	Municipio  string `xml:"xMun"`
	UF         string `xml:"uf"`
	CEP        int64  `xml:"cep"`
	Pais       string `xml:"xPais"`
	Telefone   int64  `xml:"fone"`
}

type Dest struct {
	CNPJ       int64  `xml:"cnpj"`
	Nome       string `xml:"xNome"`
	Logradouro string `xml:"lgr"`
	Numero     int64  `xml:"nro"`
	Bairro     string `xml:"bairro"`
	Municipio  string `xml:"xMun"`
	UF         string `xml:"uf"`
	CEP        int64  `xml:"cep"`
	Pais       string `xml:"xPais"`
	Telefone   int64  `xml:"fone"`
}

type Prod struct {
	CodigoID      int64   `xml:"cProd"`
	Codigo        int64   `xml:"cEAN"`
	Descricao     string  `xml:"xProd"`
	UnidMedida    string  `xml:"uCom"`
	Quantidade    int64   `xml:"qCom"`
	ValorUnitario float64 `xml:"vUnCom"`
}

// Tentativa de importar, ler e executar sobre o arquivo xml
func main() {
	resp, _ := http.Get("/home/ana/Downloads/41230910541434000152550010000012411749316397-nfe.xml")
	bytes, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	var teste Allmap
	xml.Unmarshal(bytes, &teste)

	fmt.Println(teste.Dests)

	// 	aquivoXML, erro := os.Open("/home/ana/Downloads/41230910541434000152550010000012411749316397-nfe.xml")
	// 	if erro != nil {
	// 		log.Fatalf("Erro ao abrir o arquivo xml")
	// 	} else {
	// 		fmt.Println("Arquivo aberto com sucesso")
	// 	}
	// 	defer aquivoXML.Close()
	// }

	// type destinatario struct {
}
