package modelos

import (
	"errors"
	"strings"
	"time"
)

// Publicacao representa uma publicacao de um usuario
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autoId,omitempty"`
	AutorNick uint64    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadoEm  time.Time `json:"criadaEm,omitempty"`
}

// Prepara e chama os metodos para validar e formatar a publicacao
func (publicacao *Publicacao) Preparar() error {
	if err := publicacao.validar(); err != nil {
		return err
	}
	publicacao.formatar()
	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("O Titulo é obrigatorio")
	}

	if publicacao.Conteudo == "" {
		return errors.New("O Conteudo é obrigatorio")
	}

	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
