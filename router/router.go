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
			log.Printf("Executando rota: %s comando: %q\n", rotaLocal.Nome, rotaLocal.Comando)

			id, err := model.GetChave()
			if err != nil {
				panic(err)
			}

			go func() {
				stdout, er, exitCode := rotaLocal.Funcao(rotaLocal.Comando[0], rotaLocal.Comando[1:]...)
				comando := strings.Join(rotaLocal.Comando, " ")

				controller.Resultados.Adiciona(
					id,
					model.NewSaidaComando(comando, stdout, er, exitCode),
				)
			}()

			statusCode := http.StatusOK
			c.JSON(statusCode, gin.H{
				"route":       fmt.Sprintf("/resultados/%s", id),
				"status-code": statusCode,
			})
		})
	}

	r.GET("/resultados/:id", controller.GetResultados)

	return r
}
