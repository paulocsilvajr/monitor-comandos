package view

import (
	"github.com/gin-gonic/gin"
	"github.com/paulocsilvajr/monitor-comandos/model"
)

func GetSaidaComandoJSON(comando, stdout, err string, exitCode int) gin.H {
	saidaComando := model.NewSaidaComando(comando, stdout, err, exitCode)

	return gin.H{
		comando: saidaComando.JSON(),
	}
}
