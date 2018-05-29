package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func PeliculaId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idPelicula := params["id"]
	if !bson.IsObjectIdHex(idPelicula) {
		w.WriteHeader(404)
	}
	oid := bson.ObjectIdHex(idPelicula)
	var results Pelicula
	err := Conectar.FindId(oid).One(&results)
	if err != nil {
		mostrarEror(err)
		w.WriteHeader(500)
	}
	ResponsePelicula(w, 200, results)
}
func PeliculasAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var peliculaData Pelicula
	err := decoder.Decode(&peliculaData)
	if err != nil {
		mostrarEror(err)
	}
	//limpiamos la lectura
	defer r.Body.Close()
	//guardamos en la BD
	err = Conectar.Insert(peliculaData)
	if err != nil {
		w.WriteHeader(500)
	}
	fmt.Fprintf(w, "Pelicula guardada")
	ResponsePelicula(w, 200, peliculaData)
}
func mostrarEror(err error) {
	if err != nil {
		panic(err)
	}
}
func PeliculaUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idPelicula := params["id"]
	if !bson.IsObjectIdHex(idPelicula) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(idPelicula)
	decoder := json.NewDecoder(r.Body)
	var results Pelicula
	err := decoder.Decode(&results)
	if err != nil {
		mostrarEror(err)
		w.WriteHeader(500)
		return
	}
	//update
	document := bson.M{"_id": oid}
	change := bson.M{"$set": results}
	err = Conectar.Update(document, change)

	if err != nil {
		mostrarEror(err)
		w.WriteHeader(500)
	}
	ResponsePelicula(w, 200, results)
}
func PeliculaDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idPelicula := params["id"]
	if !bson.IsObjectIdHex(idPelicula) {
		w.WriteHeader(404)
	}
	oid := bson.ObjectIdHex(idPelicula)
	// delete
	err := Conectar.RemoveId(oid)
	if err != nil {
		mostrarEror(err)
		w.WriteHeader(500)
	}
	mens := Mensage{"SUCCESS", "LA PELICULA " + idPelicula + " ha sido borrada"}
	ResponseMensage(w, 200, mens)
}
