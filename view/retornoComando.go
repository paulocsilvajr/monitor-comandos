package view

import (
	"github.com/paulocsilvajr/monitor-comandos/model"
)

func GetSaidaComandoJSON(comando, stdout, err string, exitCode int) map[string]any {
	saidaComando := model.NewSaidaComando(comando, stdout, err, exitCode)

	return saidaComando.JSON()
}
