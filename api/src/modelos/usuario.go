package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

// Preparar vai chamar os metodos para validar e formatacao
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if err := usuario.formatar(etapa); err != nil {
		return err
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("nome e obrigatorio")
	}
	if usuario.Nick == "" {
		return errors.New("nick e obrigatorio")
	}
	if usuario.Email == "" {
		return errors.New("email e obrigatorio")
	}

	if err := checkmail.ValidateFormat(usuario.Email); err != nil {
		return errors.New("email invalido")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("senha e obrigatorio")
	}
	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, err := seguranca.Hash(usuario.Senha)
		if err != nil {
			return err
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil

}
