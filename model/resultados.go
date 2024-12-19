package model

import (
	"sync"

	"github.com/google/uuid"
)

var mutex sync.Mutex

type registro map[string]SaidaComando

func GetChave() (string, error) {
	chave, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return chave.String(), nil
}

type Resultado struct {
	registros registro
}

func NewResultado() *Resultado {
	return &Resultado{make(registro)}
}

func (r *Resultado) Adiciona(chave string, s SaidaComando) {
	mutex.Lock()
	r.registros[chave] = s
	mutex.Unlock()
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

func (r *Resultado) Len() int {
	return len(r.registros)
}
