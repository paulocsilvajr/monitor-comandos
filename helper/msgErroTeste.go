package helper

import (
	"fmt"
)

func MsgErroTeste(mensagem string, valorEsperado any, valorObtido any) string {
	return fmt.Sprintf("%[1]s. Esperado: [%[2]s][%[2]T]. Obtido: [%[3]s][%[3]T]", mensagem, valorEsperado, valorObtido)
}
