package view

import "github.com/gin-gonic/gin"

func RetornaSaidaComandoJSON(rota, stdout, err string, exitCode int) gin.H {
	return gin.H{
		rota: map[string]any{
			"stdout":    stdout,
			"err":       err,
			"exit-code": exitCode,
		},
	}
}
