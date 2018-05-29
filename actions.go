package main

import "fmt"
import "net/http"
import "github.com/gorilla/mux"
import "encoding/json"
//import "strconv"
//import "log"
//import "gopkg.in/mgo.v2/bson"



func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor levantado")
}
func Hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo desde GO")
}

func Response(w http.ResponseWriter, status int , resul Peliculas) {
	
	w.Header().Set(ContentType,ApplicatioJson)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(resul)

}

func ResponsePelicula(w http.ResponseWriter, status int , resul Pelicula) {
	
	w.Header().Set(ContentType,ApplicatioJson)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(resul)

}

func ResponseMensage(w http.ResponseWriter, status int , mensage Mensage) {
	
	w.Header().Set(ContentType,ApplicatioJson)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(mensage)

}


func PelisList(w http.ResponseWriter, r *http.Request) {
	 results :=  pelisList
	//recuperar todas
//	err := Conectar.Find(nil).All(&results)

//	if err!= nil{
//		log.Println(err)
//		panic(err)
//		w.WriteHeader(500)
//		return
//	}else{
//		fmt.Println(results)
//	}
	Response(w,200,results)

}
func PeliculaId(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	idPelicula := params["id"]
	
//	if !bson.IsObjectIdHex(idPelicula){
//		w.WriteHeader(404)
//		return
//	}
//	oid := bson.ObjectIdHex(idPelicula)
//
//	var results Pelicula
//	err := Conectar.FindId(oid).One(&results)
//
//	if err!= nil {
//		mostrarEror(err)
//		w.WriteHeader(500)
//		return
//	}
	results := p1

	ResponsePelicula(w,200,results)

}
func PeliculasAdd(w http.ResponseWriter, r *http.Request) {
//	decoder := json.NewDecoder(r.Body)
//	var peliculaData Pelicula
//	err := decoder.Decode(&peliculaData)
//	if err != nil {
//		mostrarEror(err)
//	}
//	//limpiamos la lectura
//	defer r.Body.Close()
//
//	log.Println(peliculaData)
//	log.Println(peliculaData.getNombre())
//	//guardamos en la BD
//	err = Conectar.Insert(peliculaData)
//	if err!= nil {
//		w.WriteHeader(500)
//		return
//	}
	peliculaData := p2
	fmt.Fprintf(w, "Pelicula guardada")
	ResponsePelicula(w,200,peliculaData)
}
func mostrarEror(err error) {
	if err != nil {
		panic(err)
	}
}
func PeliculaUpdate(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	idPelicula := params["id"]
	
	

//	if !bson.IsObjectIdHex(idPelicula){
//		w.WriteHeader(404)
//		return
//	}
//	oid := bson.ObjectIdHex(idPelicula)
//	decoder := json.NewDecoder(r.Body)
//	var results Pelicula
//	err := decoder.Decode(&results)

//	if err!= nil {
//		mostrarEror(err)
//		w.WriteHeader(500)
//		return
//	}
	// update
//	document := bson.M{"_id": oid}
//	change := bson.M{"$set": results}
//	err = Conectar.Update(document,change)

//	if err != nil{
//		mostrarEror(err)
//		w.WriteHeader(500)
//	} 
	results := p3
	ResponsePelicula(w,200,results)

}
func PeliculaDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idPelicula := params["id"]
	
	

//	if !bson.IsObjectIdHex(idPelicula){
//		w.WriteHeader(404)
//		return
//	}
//	oid := bson.ObjectIdHex(idPelicula)
	
	// delete
	
//	err := Conectar.RemoveId(oid)

//	if err != nil{
//		mostrarEror(err)
//		w.WriteHeader(500)
//	} 

	mens := Mensage{"SUCCESS","LA PELICULA "+idPelicula+" ha sido borrada"}

	ResponseMensage(w,200,mens)

}

// creamos la variable para luego hacer inserts


type Mensage struct {
	Estado string `json:"estado"`
	Message string `json:"mensage"`
}

const (
	ApplicatioJson = "application/json"
	ContentType = "Content-Type"
	)


var p1=Pelicula{1,"Titanic","James Cameron",1990}
var p2 =Pelicula{1,"Mecawens","James Cameron",1990}
var p3 =Pelicula{1,"Terminator","James Cameron",1990}
var p4 =Pelicula{1,"Avatar","James Cameron",1990}
var pelisList = Peliculas{p1,p2,p3,p4}