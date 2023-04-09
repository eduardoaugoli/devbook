package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Escreve informacoes da requisicao no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}

}

// Verifica a existencia de cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if _, err := cookies.Ler(r); err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		proximaFuncao(w, r)
	}
}
