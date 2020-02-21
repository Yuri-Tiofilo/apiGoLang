package main

import (
	"fmt"
	"log"
	"net/http"
)

func serverConfig() {
	// configRoutes()
	fmt.Println("Servidor está rodando")
	log.Fatal(http.ListenAndServe(":3333", nil)) //DefaultServerMux
}

func main() {
	serverConfig()
}
