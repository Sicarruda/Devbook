package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/router"
)

func main() {
	config.Load()
	r := router.CreateRouter()

	fmt.Println("Rodando API")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
