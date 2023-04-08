package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	utils.CarregarTempletes()
	r := router.Gerar()
	fmt.Println("Rodando Webapp :3000!")
	log.Fatal(http.ListenAndServe(":3000", r))

}
