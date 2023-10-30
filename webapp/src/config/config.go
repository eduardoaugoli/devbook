package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//api url representa o endereço da API que será consumida
	ApiURL = ""
	//Porta representa a porta que a API será executada
	Porta = 0
	//HashKey representa a chave que será usada para assinar o token
	HashKey []byte
	//BlockKey representa a chave que será usada para criptografar o token
	BlockKey []byte
)

func Carregar() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Porta, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	ApiURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

}
