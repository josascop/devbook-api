package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

var (
	StringConexaoDB string
	PortaApi        = "9000"
)

// Carregar inicia as variáveis de ambiente da api
func Carregar() {
	varsEnv, err := godotenv.Read()
	if err != nil {
		log.Println("Erro ao iniciar variáveis de ambiente.\nImpossível continuar.\nEncerrando aplicação.")
		log.Fatal(err.Error())
	}

	p, ok := varsEnv["PORTAAPI"]
	if ok {
		PortaApi = p
	}

	StringConexaoDB = fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		varsEnv["HOSTDB"], varsEnv["PORTADB"], varsEnv["NOMEDB"], varsEnv["USERDB"], varsEnv["SENHADB"], varsEnv["SSL"])
}
