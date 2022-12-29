package respostas

import (
	"net/http"

	"github.com/go-chi/render"
)

func JSON(w http.ResponseWriter, r *http.Request, status int, dados interface{}) {
	w.WriteHeader(status)
	render.JSON(w, r, dados)
}

func Erro(w http.ResponseWriter, r *http.Request, status int, e error) {
	JSON(w, r, status, struct {
		Erro string `json:"erro"`
	}{
		Erro: e.Error(),
	})
}
