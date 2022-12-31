// Package modelos ++++++++++
package modelos

import (
	"errors"
	"net/http"
	"strings"
	"time"
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

func (u *Usuario) formatar() {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Email = strings.TrimSpace(u.Email)
	u.Nick = strings.TrimSpace(u.Nick)
}

// Bind injeta o corpo do request no struct e também faz as validações
func (u *Usuario) Bind(r *http.Request) error {
	u.formatar()

	if u.Email == "" {
		return errors.New("informe um email para o usuário")
	}
	if u.Senha == "" {
		return errors.New("informe uma senha para o usuário")
	}
	if u.Nick == "" {
		return errors.New("informe um nick para o usuário")
	}
	if u.Nome == "" {
		return errors.New("informe um nome para o usuário")
	}
	if len(u.Nome) < 3 {
		return errors.New("o nome do usuário deve ter ao menos três caracteres")
	}

	return nil
}
