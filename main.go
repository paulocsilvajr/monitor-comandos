package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func executa(comando string, argumentos ...string) string {
	cmd := exec.Command(comando, argumentos...)
	stdout, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(stdout)
}

func executa_ls() string {
	return executa("ls", "/home/paulo")
}

func executa_whoami() string {
	return executa("whoami")
}

type Rota struct {
	nome   string
	funcao func() string
}

var rotas = []Rota{
	{"whoami", executa_whoami},
	{"ls", executa_ls},
}

func main() {
	r := gin.Default()

	for _, rota := range rotas {
		// necessário declarar variável local de rota para pegar a instancia correta no loop em função anonima de r.GET
		rotaLocal := rota
		r.GET(rotaLocal.nome, func(c *gin.Context) {
			fmt.Println("Executado:", rotaLocal)
			retorno := rotaLocal.funcao()

			status := http.StatusOK
			if len(retorno) == 0 {
				status = http.StatusInternalServerError
			}

			c.JSON(status, gin.H{
				rotaLocal.nome: retorno,
			})
		})
	}

	r.Run() // 8080
}
