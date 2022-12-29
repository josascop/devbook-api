package db

import (
	"fmt"
	"josascop/calculadorago/api/src/config"
	"josascop/calculadorago/api/src/modelos"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // driver
)

var (
	db  *sqlx.DB
	err error
)

func Carregar() {
	db, err = sqlx.Connect("postgres", config.StringConexaoDB)
	if err != nil {
		log.Println(err.Error())
		fmt.Println("Impossível conectar ao banco de dados.\nEncerrando aplicação.")
		log.Fatal("Impossível conectar ao banco de dados.\nEncerrando aplicação.")
	}
	defer db.Close()

	db.MustExec(modelos.SchemaUsuario) // executa a criação da tabela de usuários
}

func Abrir() (*sqlx.DB, error) {
	db, err = sqlx.Connect("postgres", config.StringConexaoDB)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, err
}
