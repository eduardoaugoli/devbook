package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"webapp/src/respostas"
)

// CriarUsuario chama a api para cadastrar um usuario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if err != nil {
		log.Fatal(err)
	}

	url := fmt.Sprintf("%s/usuarios", "http://localhost:5000")
	response, err := http.Post(url, "application/json", bytes.NewBuffer(usuario))

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		// fmt.Println("Erro ao cadastrar o usuário:", response.Body)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		return
	}
	fmt.Println(string(body))

	respostas.JSON(w, response.StatusCode, body)
}
