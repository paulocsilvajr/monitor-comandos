package router

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/paulocsilvajr/monitor-comandos/controller"
	"github.com/paulocsilvajr/monitor-comandos/model"
)

func GetRouter(rotas []rota) *gin.Engine {
	r := gin.Default()

	for _, rota := range rotas {
		// necessário declarar variável local de rota para pegar a instancia correta no loop em função anonima de r.GET
		rotaLocal := rota
		r.GET(rotaLocal.Nome, func(c *gin.Context) {
			log.Println("Executado:", rotaLocal.Nome)

			stdout, err, exitCode := rotaLocal.Funcao(rotaLocal.Comando[0], rotaLocal.Comando[1:]...)
			comando := strings.Join(rotaLocal.Comando, " ")

			id := controller.Resultados.Adiciona(
				model.NewSaidaComando(comando, stdout, err, exitCode),
			)

			c.JSON(http.StatusOK, gin.H{
				"route": fmt.Sprintf("/resultados/%s", id),
			})
		})
	}

	r.GET("resultados/:id", controller.GetResultados)

	return r
}
