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

	if ok {
		status := http.StatusOK
		if saidaComando.ExitCode != 0 {
			status = http.StatusInternalServerError
		}

		c.JSON(status,
			view.GetSaidaComandoJSON(
				saidaComando.Comando,
				saidaComando.Stdout,
				saidaComando.Err,
				saidaComando.ExitCode,
			),
		)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "não existe resultado associado a id informada",
		})
	}

}
