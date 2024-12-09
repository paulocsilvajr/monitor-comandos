package router

import "github.com/paulocsilvajr/monitor-comandos/helper"

type rota struct {
	Nome    string
	Comando []string
	Funcao  func(string, ...string) (string, string, int)
}

func RetornaRotas(comandos []string) []rota {
	rotasComandosEParametros := helper.SeparaRotasComandosEParametros(comandos)

	var rotas = []rota{}
	for k, v := range rotasComandosEParametros {
		rotas = append(rotas, rota{k, v, helper.Executa})
	}

	return rotas
}
