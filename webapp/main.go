package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	utils.CarregarTempletes()
	fmt.Println("Rodando Webapp!")

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":3000", r))

}
