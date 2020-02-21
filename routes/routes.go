package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
	},
	Route{
		"MovieShow",
		"GET",
		"/filme/{id}",
	},
	Route{
		"MovieList",
		"GET",
		"/filmes",
	},
	Route{
		"Contact",
		"GET",
		"/contato",
	},
	Route{
		"MovieAdd",
		"POST",
		"/filme",
	},
	Route{
		"MovieUpdate",
		"PUT",
		"/filme/{id}",
	},
	Route{
		"MovieRemove",
		"DELETE",
		"/filme/{id}",
	},
}

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
