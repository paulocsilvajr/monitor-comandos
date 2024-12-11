package model

import (
	"github.com/google/uuid"
)

func GetChave() (string, error) {
	chave, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return chave.String(), nil
}

type Resultado struct {
	registros map[string]SaidaComando
}

func NewResultado() *Resultado {
	return &Resultado{make(map[string]SaidaComando)}
}

func (r *Resultado) Adiciona(chave string, s SaidaComando) {
	r.registros[chave] = s
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
