package repositorios

import (
	"database/sql"
	"fmt"
	"josascop/calculadorago/api/src/modelos"
)

const (
	queryCriar     = "INSERT INTO usuarios (nome, nick, email, senha) values ($1, $2, $3, $4);"
	queryBuscar    = "SELECT id, nome, nick, email, criadoem FROM usuarios WHERE nome LIKE $1 OR nick LIKE $1 ORDER BY nome;"
	queryBuscarID  = "SELECT id, nome, nick, email, criadoem FROM usuarios WHERE id = $1;"
	queryAtualizar = "UPDATE usuarios WHERE nome = $1, nick = $2, email = $3 WHERE id = $4;"
	queryExcluir   = "DELETE FROM usuarios WHERE id = $1;"
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

	stmt, err := repo.db.Prepare(queryBuscar)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	linhas, err := stmt.Query(nome)
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

func (repo *usuarios) BuscarID(id int64) (modelos.Usuario, error) {
	stmt, err := repo.db.Prepare(queryBuscarID)
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer stmt.Close()

	linha := stmt.QueryRow(id)
	u := modelos.Usuario{}
	if err = linha.Scan(
		&u.ID,
		&u.Nome,
		&u.Nick,
		&u.Email,
		&u.CriadoEm,
	); err != nil {
		return modelos.Usuario{}, err
	}

	return u, nil
}

func (repo *usuarios) Atualizar(id int64, mudancas modelos.Usuario) error {
	stmt, err := repo.db.Prepare(queryAtualizar)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(mudancas.Nome, mudancas.Nick, mudancas.Email, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *usuarios) Excluir(id int64) error {
	stmt, err := repo.db.Prepare(queryExcluir)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}
