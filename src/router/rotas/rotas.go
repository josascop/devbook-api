package rotas

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Rota é o tipo para guardar todas as rotas da api
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}

// Configurar adiciona todas as rotas da api ao roteador principal
func Configurar(r *chi.Mux) {
	r.Mount("/usuarios", rotasUsuarios())
}
