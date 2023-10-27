package importar

import (
	"fmt"
	"log"
	"os"
)

// Tentativa de importar, ler e executar sobre o arquivo xml
func importar() {
	aquivoXML, erro := os.Open("arquivoXML.xml")
	if erro != nil {
		log.Fatalf("Erro ao abrir o arquivo xml")
	} else {
		fmt.Println("Arquivo aberto com sucesso")
	}
	defer aquivoXML.Close()
}
