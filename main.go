package main

import (
	"fmt"
	"net/http"
)

func initialRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo Two")
}

func configRoutes() {
	http.HandleFunc("/", initialRoute)
}

func serverConfig() {
	configRoutes()
	fmt.Println("Servidor está rodando")
	http.ListenAndServe(":3333", nil) //DefaultServerMux
}

func main() {
	serverConfig()
}
