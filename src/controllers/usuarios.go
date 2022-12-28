package controllers

import (
	"josascop/calculadorago/api/src/db"
	"josascop/calculadorago/api/src/modelos"
	"josascop/calculadorago/api/src/repositorios"
	"log"
	"net/http"

	"github.com/go-chi/render"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	usuario := modelos.Usuario{}
	if err := render.Bind(r, &usuario); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	db, err := db.AbrirDB()
	if err != nil {
		db.Close()
		log.Fatalln(err)
		return
	}
	defer db.Close()

	repo := repositorios.NovoRepoUsuarios(db)

	if err := repo.Criar(usuario); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Usuário criado com sucesso"))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscando todos usuários"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscando um usuário"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualizando usuário"))
}

func ExcluirUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("excluindo usuário"))
}
