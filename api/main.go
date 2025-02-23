package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/router"
)

func main() {
	fmt.Println("Rodando API")
	r := router.CreateRouter()

	log.Fatal(http.ListenAndServe(":8080", r))

}
