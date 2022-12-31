// Package modelos ++++++++++
package modelos

import (
	"errors"
	"josascop/calculadorago/api/src/seguranca"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// SchemaUsuario é a string para criação da tabela no banco de dados pelo sqlx
const SchemaUsuario = `
CREATE TABLE IF NOT EXISTS usuarios(
	id serial primary key,
	nome varchar(50) not null,
	nick varchar(50) not null unique,
	email varchar(50) not null unique,
	senha varchar(20) not null,
	criadoem timestamp default current_timestamp
);`

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoem,omitempty" db:"criadoem"`
}

func (u *Usuario) Preparar(etapa string) error {
	if err := u.formatar(etapa); err != nil {
		return err
	}

	if err := u.validar(etapa); err != nil {
		return err
	}

	return nil
}

func (u *Usuario) formatar(etapa string) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Email = strings.TrimSpace(u.Email)
	u.Nick = strings.TrimSpace(u.Nick)

	if etapa == "inserir" {
		s, err := seguranca.Hash(u.Senha)
		if err != nil {
			return err
		}

		u.Senha = string(s)
	}

	return nil
}

// Bind injeta o corpo do request no struct e também faz as validações
func (u *Usuario) validar(etapa string) error {
	u.formatar(etapa)

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("informe um email válido")
	}
	if etapa == "inserir" && u.Senha == "" {
		return errors.New("informe uma senha válida")
	}
	if u.Nick == "" {
		return errors.New("informe um nick válido")
	}
	if u.Nome == "" {
		return errors.New("informe um nome válido")
	}
	if len(u.Nome) < 3 {
		return errors.New("o nome do usuário deve ter ao menos três caracteres")
	}

	return nil
}
