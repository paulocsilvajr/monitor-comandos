package main

import (
	"log"

	"github.com/paulocsilvajr/monitor-comandos/helper"
	"github.com/paulocsilvajr/monitor-comandos/router"
)

func main() {
	comandos, err := helper.AbreArquivoComandos()
	if err != nil {
		panic(err)
	}

	if len(comandos) > 0 {
		log.Printf("Comandos registrados em arquivo %q:\n", helper.ARQUIVO)
		for i, comando := range comandos {
			log.Printf("%d - [%s]\n", i+1, comando)
		}
		log.Println()
	} else {
		log.Fatalf("Arquivo '%s' vazio. \nAdicione comandos no arquivo citado, separados por linha, na sintaxe:\nrota comando parâmetro \"parâmetro com espaco\"\ne reexecute a API.\n", helper.ARQUIVO)
	}

	rotasComandos := router.GetRotasComandos(comandos)

	r := router.GetRouter(rotasComandos)

	r.Run() // 8080
}
