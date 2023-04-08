package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/modelos"
	"webapp/src/respostas"
)

// Fazer login utiliza e-mail e senha para autenticar
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.Erro{Erro: err.Error()})
		return
	}

	response, err := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.Erro{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.StatusCodeErro(w, response)
		return
	}

	var dadosAutenticacao modelos.DadosAutenticacao

	if err = json.NewDecoder(response.Body).Decode(&dadosAutenticacao); err != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.Erro{Erro: err.Error()})
		return
	}

	//

	respostas.JSON(w, http.StatusOK, nil)

}
