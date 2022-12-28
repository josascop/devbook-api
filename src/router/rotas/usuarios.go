package rotas

import (
	"josascop/calculadorago/api/src/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var sliceRotasUsuarios = [5]Rota{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{id}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.ExcluirUsuario,
		RequerAutenticacao: false,
	},
}

// cria o grupo de rotas de usuários para ser usado na criação do roteador principal
func rotasUsuarios() *chi.Mux {
	n := chi.NewRouter()
	n.Get("/", controllers.BuscarUsuarios)
	n.Get("/{id}", controllers.BuscarUsuario)
	n.Post("/", controllers.CriarUsuario)
	n.Put("/{id}", controllers.AtualizarUsuario)
	n.Delete("/{id}", controllers.ExcluirUsuario)

	return n
}
