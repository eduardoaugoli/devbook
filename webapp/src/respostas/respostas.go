package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// Representa a resposta de erro da api
type Erro struct {
	Erro string `json:"erro"`
}

// JSON retorna uma respostas em formato JSON para a requisicao
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(dados); err != nil {
		log.Fatal(err)
	}
}

// trata erro 400 ou superior
func StatusCodeErro(w http.ResponseWriter, r *http.Response) {
	var err Erro

	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, r.StatusCode, err)
}
