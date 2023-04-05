package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// Carregar tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplete(w, "login.html", nil)

}
