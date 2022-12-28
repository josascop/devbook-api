package main

import (
	"fmt"
	"josascop/calculadorago/api/src/config"
	"josascop/calculadorago/api/src/db"
	"josascop/calculadorago/api/src/router"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	log.Println("Variáveis de ambiente carregadas.")
	db.Carregar()
	log.Println("Banco de dados configurado.")

	router := router.Gerar()

	log.Printf("Aplicação iniciada na porta %s.", config.PortaApi)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.PortaApi), router))
}
