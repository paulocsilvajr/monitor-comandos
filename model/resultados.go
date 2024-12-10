package model

import (
	"github.com/google/uuid"
)

type Resultado struct {
	registros map[string]SaidaComando
}

func NewResultado() *Resultado {
	return &Resultado{make(map[string]SaidaComando)}
}

func (r *Resultado) Adiciona(s SaidaComando) string {
	chave, err := uuid.NewRandom()
	if err != nil {
		return ""
	}

	r.registros[chave.String()] = s
	return chave.String()
}

func (r *Resultado) Get(chave string) (SaidaComando, bool) {
	if v, ok := r.registros[chave]; ok {
		return v, ok
	} else {
		return SaidaComando{}, ok
	}
}

func (r *Resultado) Remove(id string) {
	delete(r.registros, id)
}
