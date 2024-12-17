package helper

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const ARQUIVO = "comandos"

func AbreArquivoComandos() ([]string, error) {
	caminhoArquivo, err := getCaminhoArquivoComandos()
	if err != nil {
		return nil, err
	}
	arquivo, err := os.OpenFile(caminhoArquivo, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer arquivo.Close()
	log.Printf("Arquivo de comandos em %q", caminhoArquivo)

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

func getCaminhoArquivoComandos() (string, error) {
	diretorioExecutavel, err := getDiretorioAbs()
	if err != nil {
		return "", err
	}

	caminhoArquivo := path.Join(diretorioExecutavel, ARQUIVO)
	return caminhoArquivo, nil
}

func getDiretorioAbs() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
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
