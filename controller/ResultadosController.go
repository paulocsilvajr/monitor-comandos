package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paulocsilvajr/monitor-comandos/model"
	"github.com/paulocsilvajr/monitor-comandos/view"
)

var Resultados = model.NewResultado()

func GetResultados(c *gin.Context) {
	id := c.Param("id")

	saidaComando, ok := Resultados.Get(id)
	statusCode := http.StatusOK
	if ok {
		if saidaComando.ExitCode != 0 {
			statusCode = http.StatusInternalServerError
		}

		c.JSON(statusCode, gin.H{
			saidaComando.Comando: view.GetSaidaComandoJSON(
				saidaComando.Comando,
				saidaComando.Stdout,
				saidaComando.Err,
				saidaComando.ExitCode,
			),
			"status-code": statusCode,
		})

		Resultados.Remove(id)
	} else {
		statusCode = http.StatusNotFound
		c.JSON(statusCode, gin.H{
			"message":     "não existe resultado associado a id informada",
			"status-code": statusCode,
		})
	}

}
