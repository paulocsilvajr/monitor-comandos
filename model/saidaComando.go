package model

// type SaidaComando map[string]any
type SaidaComando struct {
	Comando  string
	Stdout   string
	Err      string
	ExitCode int
}

func NewSaidaComando(comando, stdout, err string, exitCode int) SaidaComando {
	return SaidaComando{comando, stdout, err, exitCode}
}

func (s SaidaComando) JSON() map[string]any {
	mapSaidaComando := make(map[string]any)
	mapSaidaComando["stdout"] = s.Stdout
	mapSaidaComando["err"] = s.Err
	mapSaidaComando["exit-code"] = s.ExitCode

	return mapSaidaComando
}
