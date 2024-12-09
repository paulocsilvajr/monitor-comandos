package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paulocsilvajr/monitor-comandos/view"
)

func RetornaRouter(rotas []rota) *gin.Engine {
	r := gin.Default()

	for _, rota := range rotas {
		// necessário declarar variável local de rota para pegar a instancia correta no loop em função anonima de r.GET
		rotaLocal := rota
		r.GET(rotaLocal.Nome, func(c *gin.Context) {
			fmt.Println("Executado:", rotaLocal.Nome)
			stdout, err, exitCode := rotaLocal.Funcao(rotaLocal.Comando[0], rotaLocal.Comando[1:]...)

			status := http.StatusOK
			if exitCode != 0 {
				status = http.StatusInternalServerError
			}

			c.JSON(status,
				view.RetornaSaidaComandoJSON(
					rotaLocal.Nome,
					stdout,
					err,
					exitCode,
				),
			)
		})
	}

	return r
}
