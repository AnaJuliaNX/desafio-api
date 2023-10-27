package importar

import (
	"fmt"
	"log"
	"os"
)

// Tentativa de importar, ler e executar sobre o arquivo xml
func main() {
	aquivoXML, erro := os.Open("/home/ana/Downloads/41230910541434000152550010000012411749316397-nfe.xml")
	if erro != nil {
		log.Fatalf("Erro ao abrir o arquivo xml")
	} else {
		fmt.Println("Arquivo aberto com sucesso")
	}
	defer aquivoXML.Close()
}

// type destinatario struct {
// }
