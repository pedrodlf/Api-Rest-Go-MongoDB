package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func main() {

	router := NewRouter()

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)

}

var Conectar = GetSession().DB("atSistemas").C("peliculas")

func GetSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		mostrarEror(err)
	}
	return session
}
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).Handler(route.HandleFunc)
	}
	return router

}
