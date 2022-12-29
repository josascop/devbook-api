package controllers

import (
	"josascop/calculadorago/api/src/db"
	"josascop/calculadorago/api/src/modelos"
	"josascop/calculadorago/api/src/repositorios"
	"josascop/calculadorago/api/src/respostas"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	usuario := modelos.Usuario{}
	if err := render.Bind(r, &usuario); err != nil {
		log.Println(err)
		respostas.Erro(w, r, http.StatusBadRequest, err)
		return
	}

	db, err := db.Abrir()
	if err != nil {
		db.Close()
		log.Println(err)
		respostas.Erro(w, r, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositorios.NovoRepoUsuarios(db)
	if err := repo.Criar(usuario); err != nil {
		log.Println(err)
		respostas.Erro(w, r, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, r, http.StatusCreated, "Usuário criado com sucesso")
}

// /usuarios?busca=junin
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nome := strings.ToLower(chi.URLParam(r, "busca"))

	db, err := db.Abrir()
	if err != nil {
		db.Close()
		log.Println(err)
		respostas.Erro(w, r, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositorios.NovoRepoUsuarios(db)
	usuarios, err := repo.Buscar(nome)
	if err != nil {
		log.Println(err)
		respostas.Erro(w, r, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, r, http.StatusOK, usuarios)
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
