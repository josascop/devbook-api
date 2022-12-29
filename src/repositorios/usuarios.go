package repositorios

import (
	"fmt"
	"josascop/calculadorago/api/src/modelos"

	"github.com/jmoiron/sqlx"
)

const (
	queryCriar  = "INSERT INTO usuarios (nome, nick, email, senha) values ($1, $2, $3, $4);"
	queryBuscar = "SELECT (id, nome, nick, email, criadoem) FROM usuarios WHERE nome LIKE $1 OR nick LIKE $1 ORDER BY nome;"
)

type usuarios struct {
	db *sqlx.DB
}

func NovoRepoUsuarios(db *sqlx.DB) *usuarios {
	return &usuarios{db}
}

func (repo *usuarios) Criar(u modelos.Usuario) error {
	_, err := repo.db.Exec(queryCriar, u.Nome, u.Nick, u.Email, u.Senha)
	if err != nil {
		return err
	}

	return nil
}

// Buscar busca todos os usuários que correspondem a nome em Nome ou Nick
func (repo *usuarios) Buscar(nome string) ([]modelos.Usuario, error) {
	nome = fmt.Sprintf("%%%s%%", nome)

	resp := []modelos.Usuario{}
	// us := modelos.Usuario{}
	// linhas, err := repo.db.Queryx(queryBuscar, nome)
	// if err != nil {
	// 	return nil, err
	// }
	// for linhas.Next() {
	// 	if err = linhas.StructScan(&us); err != nil {
	// 		return nil, err
	// 	}
	// 	resp = append(resp, us)
	// }
	if err := repo.db.Select(&resp, queryBuscar, nome); err != nil {
		return nil, err
	}

	return resp, nil
}
