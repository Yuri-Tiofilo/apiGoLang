package main

import (
	"fmt"
	"log"
	"net/http"
)

func serverConfig() {
	teste := NewRouter()
	fmt.Println(teste)
	configRoutes()
	fmt.Println("Servidor está rodando")
	log.Fatal(http.ListenAndServe(":3333", nil)) //DefaultServerMux
}

func main() {
	serverConfig()
}
