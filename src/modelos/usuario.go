package modelos

import (
	"net/http"
	"time"
)

const SchemaUsuario = `
CREATE TABLE IF NOT EXISTS usuarios(
	id serial primary key,
	nome varchar(50) not null,
	nick varchar(50) not null unique,
	email varchar(50) not null unique,
	senha varchar(20) not null,
	criadoEm timestamp default current_timestamp
);`

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Bind injeta o corpo do request no struct e também faz as validações
func (u Usuario) Bind(r *http.Request) error {
	return nil
}
