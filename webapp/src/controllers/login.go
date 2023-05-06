package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
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

	url := fmt.Sprintf("%s/login", config.APIURL)

	response, err := http.Post(url, "application/json", bytes.NewBuffer(usuario))
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

	if err = cookies.Salvar(w, dadosAutenticacao.ID, dadosAutenticacao.Token); err != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.Erro{Erro: err.Error()})
		return
	}

	respostas.JSON(w, http.StatusOK, nil)

}
