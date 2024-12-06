package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"syscall"

	"github.com/gin-gonic/gin"
)

func executa(comando string, argumentos ...string) (stdout string, stderr string, exitCode int) {
	cmd := exec.Command(comando, argumentos...)
	combinedOutput, err := cmd.CombinedOutput()
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return string(combinedOutput), err.Error(), status.ExitStatus()
			}
		}
		return "", err.Error(), -1
	}
	return string(combinedOutput), "", 0
}

func executa_ls() (string, string, int) {
	return executa("ls", "/home/paulo")
}

func executa_whoami() (string, string, int) {
	return executa("whoami")
}

func executa_naoexiste() (string, string, int) {
	return executa("naoexiste")
}

func executa_ls2() (string, string, int) {
	return executa("ls", "2")
}

func execute_catwords() (string, string, int) {
	return executa("cat", "/etc/dictionaries-common/words")
}

type Rota struct {
	nome   string
	funcao func() (string, string, int)
}

var rotas = []Rota{
	{"whoami", executa_whoami},
	{"ls", executa_ls},
	{"naoexiste", executa_naoexiste},
	{"ls2", executa_ls2},
	{"catwords", execute_catwords},
}

func main() {
	r := gin.Default()

	for _, rota := range rotas {
		// necessário declarar variável local de rota para pegar a instancia correta no loop em função anonima de r.GET
		rotaLocal := rota
		r.GET(rotaLocal.nome, func(c *gin.Context) {
			fmt.Println("Executado:", rotaLocal)
			stdout, err, exitCode := rotaLocal.funcao()

			status := http.StatusOK
			if exitCode != 0 {
				status = http.StatusInternalServerError
			}

			c.JSON(status, gin.H{
				rotaLocal.nome: map[string]any{
					"stdout":    stdout,
					"err":       err,
					"exit-code": exitCode,
				},
			})
		})
	}

	r.Run() // 8080
}
