package model

import (
	"testing"

	"github.com/paulocsilvajr/monitor-comandos/helper"
)

func TestAdiciona(t *testing.T) {
	resultado := NewResultado()
	comando := "ls"
	stdout := "teste comando"
	err := ""
	exitCode := 0
	saidaComando := NewSaidaComando(comando, stdout, err, exitCode)

	chave := resultado.Adiciona(saidaComando)

	if _, ok := resultado.registros[chave]; !ok {
		t.Error(helper.MsgErroTeste("chave não existe", chave, ""))
	}

	if valor, ok := resultado.registros[chave]; ok {
		if valor.Stdout != stdout {
			t.Error(helper.MsgErroTeste("valor stdout errado", stdout, valor.Stdout))
		}
		if valor.Err != err {
			t.Error(helper.MsgErroTeste("valor err errado", err, valor.Err))
		}
		if valor.ExitCode != exitCode {
			t.Error(helper.MsgErroTeste("valor exit-code errado", exitCode, valor.ExitCode))
		}
	}
}

func TestAdicionaVariosRegistros(t *testing.T) {
	resultado := NewResultado()

	stdout01 := "saída comando 01"
	comando01 := "ls"
	err01 := ""
	exitCode01 := 0
	chave01 := resultado.Adiciona(NewSaidaComando(comando01, stdout01, err01, exitCode01))

	stdout02 := ""
	comando02 := "ls"
	err02 := "erro comando 02"
	exitCode02 := 2
	chave02 := resultado.Adiciona(NewSaidaComando(comando02, stdout02, err02, exitCode02))

	stdout03 := "saída comando 03"
	comando03 := "ls"
	err03 := ""
	exitCode03 := 0
	chave03 := resultado.Adiciona(NewSaidaComando(comando03, stdout03, err03, exitCode03))

	if valor, ok := resultado.Get(chave01); ok {
		if valor.Stdout != stdout01 {
			t.Error(helper.MsgErroTeste("valor stdout01 errado", stdout01, valor.Stdout))
		}
		if valor.Err != err01 {
			t.Error(helper.MsgErroTeste("valor err01 errado", err01, valor.Err))
		}
		if valor.ExitCode != exitCode01 {
			t.Error(helper.MsgErroTeste("valor exit-code01 errado", exitCode01, valor.ExitCode))
		}
	}

	if valor, ok := resultado.Get(chave02); ok {
		if valor.Stdout != stdout02 {
			t.Error(helper.MsgErroTeste("valor stdout02 errado", stdout02, valor.Stdout))
		}
		if valor.Err != err02 {
			t.Error(helper.MsgErroTeste("valor err02 errado", err02, valor.Err))
		}
		if valor.ExitCode != exitCode02 {
			t.Error(helper.MsgErroTeste("valor exit-code02 errado", exitCode02, valor.ExitCode))
		}
	}

	if valor, ok := resultado.Get(chave03); ok {
		if valor.Stdout != stdout03 {
			t.Error(helper.MsgErroTeste("valor stdout03 errado", stdout03, valor.Stdout))
		}
		if valor.Err != err03 {
			t.Error(helper.MsgErroTeste("valor err03 errado", err03, valor.Err))
		}
		if valor.ExitCode != exitCode03 {
			t.Error(helper.MsgErroTeste("valor exit-code03 errado", exitCode03, valor.ExitCode))
		}
	}
}

func TestGetChaveValida(t *testing.T) {
	resultado := NewResultado()
	comando := "ls"
	stdout := "teste comando"
	err := "teste erro"
	exitCode := 42

	chave := resultado.Adiciona(NewSaidaComando(comando, stdout, err, exitCode))

	valor, ok := resultado.Get(chave)
	if !ok {
		t.Error(helper.MsgErroTeste("não encontrou chave válida", "", valor))
	}
	if valor.Stdout != stdout {
		t.Error(helper.MsgErroTeste("valor stdout errado", stdout, valor.Stdout))
	}
	if valor.Err != err {
		t.Error(helper.MsgErroTeste("valor err errado", err, valor.Err))
	}
	if valor.ExitCode != exitCode {
		t.Error(helper.MsgErroTeste("valor exit-code errado", exitCode, valor.ExitCode))
	}
}

func TestGetChaveInvalida(t *testing.T) {
	resultado := NewResultado()
	comando := "ls"
	stdout := "teste comando"
	err := "teste erro"
	exitCode := 42

	resultado.Adiciona(NewSaidaComando(comando, stdout, err, exitCode))

	chaveInvalida := "ch4v3Qu3N403x15t3"
	valor, ok := resultado.Get(chaveInvalida)
	if ok {
		t.Error(helper.MsgErroTeste("encontrou valor com chave inválida", false, ok))
	}
	if valor.Stdout == stdout {
		t.Error(helper.MsgErroTeste("valor stdout errado", "", valor.Stdout))
	}
	if valor.Err == err {
		t.Error(helper.MsgErroTeste("valor err errado", err, valor.Err))
	}
	if valor.ExitCode == exitCode {
		t.Error(helper.MsgErroTeste("valor exit-code errado", exitCode, valor.ExitCode))
	}
}

func TestRemove(t *testing.T) {
	resultado := NewResultado()
	comando := "ls"
	stdout := "saida comando com sucesso"
	err := ""
	exitCode := 0

	chave := resultado.Adiciona(NewSaidaComando(comando, stdout, err, exitCode))

	resultado.Remove(chave)

	if valor, ok := resultado.Get(chave); ok {
		t.Error(helper.MsgErroTeste("retornou valor para chave removida", nil, valor))
	}
}

func TestRemoveVariosRegistros(t *testing.T) {
	resultado := NewResultado()
	comando01 := "ls"
	stdout01 := "saida comando com sucesso"
	err01 := ""
	exitCode01 := 0
	chave01 := resultado.Adiciona(NewSaidaComando(comando01, stdout01, err01, exitCode01))

	comando02 := "rs"
	stdout02 := ""
	err02 := "erro comando"
	exitCode02 := -1
	chave02 := resultado.Adiciona(NewSaidaComando(comando02, stdout02, err02, exitCode02))

	comando03 := "ts"
	stdout03 := "saida comando 03"
	err03 := "erro comando 03"
	exitCode03 := 2
	chave03 := resultado.Adiciona(NewSaidaComando(comando03, stdout03, err03, exitCode03))

	resultado.Remove(chave01)
	if valor, ok := resultado.Get(chave01); ok {
		t.Error(helper.MsgErroTeste("retornou valor para chave01 removida", nil, valor))
	}

	resultado.Remove(chave02)
	if valor, ok := resultado.Get(chave02); ok {
		t.Error(helper.MsgErroTeste("retornou valor para chave02 removida", nil, valor))
	}

	if _, ok := resultado.Get(chave03); !ok {
		t.Error(helper.MsgErroTeste("não encontrou valor com chave03 válida", true, false))
	}

	resultado.Remove(chave03)
	if valor, ok := resultado.Get(chave03); ok {
		t.Error(helper.MsgErroTeste("retornou valor para chave03 removida", nil, valor))
	}

}
