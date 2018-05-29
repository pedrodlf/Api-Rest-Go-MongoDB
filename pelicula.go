package main

type Pelicula struct {
	Id       int    // `json:"id"`
	Nombre   string `json:"name"`
	Director string `json:"director"`
	Anio     int    `json:"year"`
}



type Peliculas []Pelicula

func (this *Pelicula) getNombre() string{
	return this.Nombre
}

func (this *Pelicula) getDirector() string{
	return this.Director
}

func (este *Pelicula) setNombre(nombre string) {
	este.Nombre = nombre
}