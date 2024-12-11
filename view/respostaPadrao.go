package view

import "github.com/gin-gonic/gin"

func resposta(statusCode int, valores map[string]any) gin.H {
	resposta := make(gin.H)
	resposta["status-code"] = statusCode

	for k, v := range valores {
		resposta[k] = v
	}

	return resposta
}

func RespostaJSON(c *gin.Context, statusCode int, valores map[string]any) {
	c.JSON(statusCode, resposta(
		statusCode,
		valores),
	)
}
