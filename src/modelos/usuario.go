// Package modelos ++++++++++
package modelos

import (
	"errors"
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
	u.formatar()

	if err := u.validar(etapa); err != nil {
		return err
	}

	return nil
}

func (u *Usuario) formatar() {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Email = strings.TrimSpace(u.Email)
	u.Nick = strings.TrimSpace(u.Nick)
}

// Bind injeta o corpo do request no struct e também faz as validações
func (u *Usuario) validar(etapa string) error {
	u.formatar()

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
