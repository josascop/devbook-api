package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"josascop/calculadorago/api/src/db"
	"josascop/calculadorago/api/src/modelos"
	"josascop/calculadorago/api/src/repositorios"
	"josascop/calculadorago/api/src/respostas"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	usuario := modelos.Usuario{}
	if err := json.Unmarshal(corpo, &usuario); err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Preparar("inserir"); err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Abrir()
	if err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositorios.NovoRepoUsuarios(db)
	if err := repo.Criar(usuario); err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, "Usuário criado com sucesso")
}

// /usuarios?busca=junin
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("busca")

	db, err := db.Abrir()
	if err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositorios.NovoRepoUsuarios(db)
	usuarios, err := repo.Buscar(nome)
	if err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Abrir()
	if err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositorios.NovoRepoUsuarios(db)
	usuario, err := repo.BuscarID(id)
	if err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusNotFound, errors.New("usuário não encontrado"))
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	corpo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	mudancas := modelos.Usuario{}
	if err = json.Unmarshal(corpo, &mudancas); err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = mudancas.Preparar("editar"); err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Abrir()
	if err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositorios.NovoRepoUsuarios(db)
	if err = repo.Atualizar(id, mudancas); err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func ExcluirUsuario(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Abrir()
	if err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositorios.NovoRepoUsuarios(db)
	if err = repo.Excluir(id); err != nil {
		log.Println(err)
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
