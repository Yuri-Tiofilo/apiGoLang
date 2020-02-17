package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type livro struct {
	Id     int
	Title  string
	Author string
}

var Livros []livro = []livro{
	livro{
		Id:     1,
		Title:  "A vida",
		Author: "Yuri",
	},
	livro{
		Id:     2,
		Title:  "A vida",
		Author: "Yuri",
	},
	livro{
		Id:     3,
		Title:  "A vida",
		Author: "Yuri",
	},
	livro{
		Id:     4,
		Title:  "A vida",
		Author: "Yuri",
	},
}

func initialRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo Two")
}

func listBooks(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(Livros)
}

func configRoutes() {
	http.HandleFunc("/", initialRoute)
	http.HandleFunc("/books", listBooks)
}

func serverConfig() {
	configRoutes()
	fmt.Println("Servidor está rodando")
	log.Fatal(http.ListenAndServe(":3333", nil)) //DefaultServerMux
}

func main() {
	routes()
	serverConfig()
}
