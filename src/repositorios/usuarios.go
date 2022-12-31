package repositorios

import (
	"database/sql"
	"fmt"
	"josascop/calculadorago/api/src/modelos"
)

const (
	queryCriar  = "INSERT INTO usuarios (nome, nick, email, senha) values ($1, $2, $3, $4);"
	queryBuscar = "SELECT id, nome, nick, email, criadoem FROM usuarios WHERE nome LIKE $1 OR nick LIKE $1 ORDER BY nome;"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepoUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (repo *usuarios) Criar(u modelos.Usuario) error {
	stmt, err := repo.db.Prepare(queryCriar)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Nome, u.Nick, u.Email, u.Senha)
	if err != nil {
		return err
	}

	return nil
}

// Buscar busca todos os usuários que correspondem a nome em Nome ou Nick
func (repo *usuarios) Buscar(nome string) ([]modelos.Usuario, error) {
	nome = fmt.Sprintf("%%%s%%", nome)

	linhas, err := repo.db.Query(queryBuscar, nome)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	resp := []modelos.Usuario{}
	for linhas.Next() {
		us := modelos.Usuario{}
		if err = linhas.Scan(
			&us.ID,
			&us.Nome,
			&us.Nick,
			&us.Email,
			&us.CriadoEm,
		); err != nil {
			return nil, err
		}

		resp = append(resp, us)
	}

	return resp, nil
}
