package router

import (
	"webapp/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gera retorna um router com todas as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)
}
