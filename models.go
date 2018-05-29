package main

import (
	"encoding/json"
	"net/http"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{"PeliculasList", "GET", "/peliculas/{id}", PeliculaId},
	Route{"PeliculasUpdate", "PUT", "/peliculas/{id}", PeliculaUpdate},
	Route{"PeliculasDelete", "DELETE", "/peliculas/{id}", PeliculaDelete},
	Route{"PeliculasAdd", "POST", "/peliculas", PeliculasAdd}}

// creamos la variable para luego hacer inserts

type Mensage struct {
	Estado  string `json:"estado"`
	Message string `json:"mensage"`
}

const (
	ApplicatioJson = "application/json"
	ContentType    = "Content-Type"
)

type Pelicula struct {
	Id       int    // `json:"id"`
	Nombre   string `json:"name"`
	Director string `json:"director"`
	Anio     int    `json:"year"`
}

type Peliculas []Pelicula

func ResponsePelicula(w http.ResponseWriter, status int, resul Pelicula) {

	w.Header().Set(ContentType, ApplicatioJson)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(resul)

}
func ResponseMensage(w http.ResponseWriter, status int, mensage Mensage) {

	w.Header().Set(ContentType, ApplicatioJson)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(mensage)

}
