package helper

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const ARQUIVO = "comandos"

func AbreArquivoComandos() ([]string, error) {
	arquivo, err := os.OpenFile(ARQUIVO, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	var comandos []string
	for scanner.Scan() {
		linha := scanner.Text()
		comandos = append(comandos, linha)
	}

	return comandos, nil
}

func SeparaRotasComandosEParametros(comandosEmArquivo []string) map[string][]string {
	comandosEmTokens := converteArquivoEmTokens(comandosEmArquivo)

	rotasComandosEParametros := make(map[string][]string)
	for _, tokens := range comandosEmTokens {
		converteTokensEmMapRotas(tokens, rotasComandosEParametros)
	}

	return rotasComandosEParametros
}

func converteTokensEmMapRotas(tokens []string, rotasComandosEParametros map[string][]string) {
	var comandos []string
	var nomeComEspacos string
	var nomeComposto bool

	chave := tokens[0]
	for _, token := range tokens[1:] {
		if strings.HasPrefix(token, "\"") {
			nomeComEspacos += token
			nomeComposto = true
		} else if nomeComposto {
			if strings.HasSuffix(token, "\"") {
				nomeComEspacos += fmt.Sprintf(" %s", token)
				comandos = append(comandos, nomeComEspacos)
				nomeComEspacos = ""
				nomeComposto = false
			} else {
				nomeComEspacos += fmt.Sprintf(" %s", token)
			}
		} else {
			comandos = append(comandos, token)
		}
	}
	rotasComandosEParametros[chave] = comandos
}

func converteArquivoEmTokens(comandosEmArquivo []string) [][]string {
	var comandosEmTokens [][]string
	for _, comando := range comandosEmArquivo {
		tokens := strings.Fields(comando)
		comandosEmTokens = append(comandosEmTokens, tokens)
	}
	return comandosEmTokens
}
