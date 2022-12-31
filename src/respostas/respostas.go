package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, dados interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if dados != nil {
		if err := json.NewEncoder(w).Encode(dados); err != nil {
			log.Fatal(err)
		}
	}
}

func Erro(w http.ResponseWriter, status int, e error) {
	JSON(w, status, struct {
		Erro string `json:"erro"`
	}{
		Erro: e.Error(),
	})
}
