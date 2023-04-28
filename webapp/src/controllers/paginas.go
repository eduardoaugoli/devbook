package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// Carregar tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplete(w, "login.html", nil)

}

// Carrega pagina de cadastro de usuario
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplete(w, "cadastro.html", nil)
}

// Carrega pagina principal da aplicacao
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, err := requisicoes.ExecRequestAuth(r, http.MethodGet, url, nil)
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.Erro{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.StatusCodeErro(w, response)
		return
	}

	var publicacoes []modelos.Publicacao
	if err = json.NewDecoder(response.Body).Decode(&publicacoes); err != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.Erro{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecutarTemplete(w, "home.html", struct {
		Publicacoes []modelos.Publicacao
		UsuarioID   uint64
	}{
		Publicacoes: publicacoes,
		UsuarioID:   usuarioID,
	})
}

func CarregaEdicaoPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.Erro{Erro: err.Error()})
		return
	}

	url := fmt.Sprint("%s/publicacoes/%d", config.APIURL, publicacaoID)
	response, err := requisicoes.ExecRequestAuth(r, http.MethodGet, url, nil)
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.Erro{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.StatusCodeErro(w, response)
		return
	}

	var publicacao modelos.Publicacao
	if err = json.NewDecoder(response.Body).Decode(&publicacao); err != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.Erro{Erro: err.Error()})
		return
	}

	utils.ExecutarTemplete(w, "atualizar-publicacao.html", publicacao)
}
