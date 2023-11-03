package importar

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

// A struct NFe só da certo se estiver implementada nessa struct
type NfeProc struct {
	NFe NFe `xml:"NFe"`
}

// A struct InfNFe só da certo se estiver implementada nessa struct
type NFe struct {
	InfNFe InfNFe `xml:"infNFe"`
}

// As structs Emit, Dest e Det só da certo se estiverem implementadas nessa struct
type InfNFe struct {
	Emit Emit `xml:"emit"`
	Dest Dest `xml:"dest"`
	Det  Det  `xml:"det"`
}
type Emit struct {
	CNPJ       string `xml:"CNPJ"`
	Nome       string `xml:"xNome"`
	Logradouro string `xml:"enderEmit>Lgr"`
	Numero     string `xml:"enderEmit>nro"`
	Bairro     string `xml:"enderEmit>xBairro"`
	Municipio  string `xml:"enderEmit>xMun"`
	UF         string `xml:"enderEmit>UF"`
	CEP        string `xml:"enderEmit>CEP"`
	Pais       string `xml:"enderEmit>Pais"`
	Telefone   string `xml:"enderEmit>fone"`
}

type Dest struct {
	CNPJ       string `xml:"CNPJ"`
	Logradouro string `xml:"enderDest>xLgr"`
	Numero     string `xml:"enderDest>nro"`
	Bairro     string `xml:"enderDest>xBairro"`
	Municipio  string `xml:"enderDest>xMun"`
	UF         string `xml:"enderDest>UF"`
	CEP        string `xml:"enderDest>CEP"`
	Pais       string `xml:"enderDest>xPais"`
	Telefone   string `xml:"enderDest>fone"`
}

type Det struct {
	Nitem string `xml:"nItem,attr"`
	Prod  []Prod `xml:"prod"`
}

type Prod struct {
	Codigo        string `xml:"cEAN"`
	Descricao     string `xml:"xProd"`
	UniMedida     string `xml:"uCom"`
	Quantidade    string `xml:"qCom"`
	ValorUnitario string `xml:"vUnCom"`
	Frete         string `xml:"vFrete"` //somatório do frete dos produtos
	Seguro        string `xml:"vSeg"`   //somatório do valor do seguro de todos os produtos
	Desconto      string `xml:"vDesc"`  //somátorio dos descontos de todos os produtos
	Outro         string `xml:"vOutro"` //outros valores que podem constar também
}

func Importar() {
	arquivo, erro := os.Open("/home/ana/Downloads/41230910541434000152550010000012411749316397-nfe.xml")
	if erro != nil {
		log.Fatalf("Erro ao abrir o arquivo XML")
	}
	defer arquivo.Close()

	lerarquivo, erro1 := io.ReadAll(arquivo)
	if erro1 != nil {
		log.Fatalf("Erro ao ler o conteúdo do arquivo")
	}

	var XML1 NfeProc
	erro2 := xml.Unmarshal(lerarquivo, &XML1)
	if erro2 != nil {
		log.Fatalf("Erro ao fazer o Unmarshal")
	}

	fmt.Println(XML1.NFe.InfNFe.Emit)
	fmt.Println(XML1.NFe.InfNFe.Dest)
	fmt.Println(XML1.NFe.InfNFe.Det.Prod)
}
