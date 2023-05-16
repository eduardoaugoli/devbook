package router

import "github.com/gorilla/mux"

// Gerar retorna um router com todas as configurações
func Gerar() *mux.Router {
	return mux.NewRouter()
}
