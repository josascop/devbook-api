package db

import (
	"database/sql"
	"fmt"
	"josascop/calculadorago/api/src/config"
	"josascop/calculadorago/api/src/modelos"
	"log"

	_ "github.com/lib/pq" // driver
)

var (
	db  *sql.DB
	err error
)

func Carregar() {
	db, err = sql.Open("postgres", config.StringConexaoDB)
	if err != nil {
		log.Println(err.Error())
		fmt.Println("Impossível conectar ao banco de dados.\nEncerrando aplicação.")
		log.Fatal("Impossível conectar ao banco de dados.\nEncerrando aplicação.")
	}
	defer db.Close()

	_, err = db.Exec(modelos.SchemaUsuario) // executa a criação da tabela de usuários
	if err != nil {
		log.Println(err.Error())
		fmt.Println("Impossível conectar ao banco de dados.\nEncerrando aplicação.")
		log.Fatal("Impossível conectar ao banco de dados.\nEncerrando aplicação.")
	}
}

func Abrir() (*sql.DB, error) {
	db, err = sql.Open("postgres", config.StringConexaoDB)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, err
}
