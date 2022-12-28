package repositorios

import (
	"josascop/calculadorago/api/src/modelos"

	"github.com/jmoiron/sqlx"
)

const (
	queryInserir = "INSERT INTO usuarios (nome, nick, email, senha) values ($1, $2, $3, $4);"
)

type usuarios struct {
	db *sqlx.DB
}

func NovoRepoUsuarios(db *sqlx.DB) *usuarios {
	return &usuarios{db}
}

func (repo *usuarios) Criar(u modelos.Usuario) error {
	_, err := repo.db.Exec(queryInserir, u.Nome, u.Nick, u.Email, u.Senha)
	if err != nil {
		return err
	}

	return nil
}
