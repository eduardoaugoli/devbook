package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//representa a url para comunicacao com a API
	APIURL = ""
	//Porta Aplicacao web esta rodando
	Porta = 0
	//HashKey é utilizado para autenticar o cookie
	HashKey []byte
	//BlockKey e utilizada para criptografar os dados do cookie
	BlockKey []byte
)

// Carregar inicializa as variaveis de ambiente
func Carregar() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Porta, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

}
