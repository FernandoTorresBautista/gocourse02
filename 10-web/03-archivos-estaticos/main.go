package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Funciones
func Saludar(nombre string) string {
	return "Hola " + nombre + " desde una función"
}

// structs

type Usuarios struct {
	Name   string
	Age    int
	Activo bool
	Admin  bool
	Cursos []Curso
}

type Curso struct {
	Nombre string
}

// importar todos los archivos html
var temp = template.Must(template.New("T").ParseGlob("templates/**/*.html"))
var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

// handlers

func Index(w http.ResponseWriter, r *http.Request) {
	c1 := Curso{"Go"}
	c2 := Curso{"Java"}
	c3 := Curso{"Python"}
	c4 := Curso{"JS"}

	cursos := []Curso{c1, c2, c3, c4}
	user := Usuarios{"Fer", 26, true, false, cursos}

	// usar el temp con todas las variables
	if err := temp.ExecuteTemplate(w, "index.html", user); err != nil {
		log.Fatal(err)
	}
}

// register, change the render to a function to no repeat the same code
func Registro(w http.ResponseWriter, r *http.Request) {
	// if err := temp2.ExecuteTemplate(w, "registro.html", nil); err != nil {
	// 	log.Fatal(err)
	// }
	renderTemplate(w, "registro.html", nil)
}

// handler error
func hadlerError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	if err := errorTemplate.Execute(w, nil); err != nil {
		panic("Error al recuperar el template de error " + err.Error())
	}
}

// funcion de render template
func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	if err := temp.ExecuteTemplate(w, name, data); err != nil {
		// http.Error(w, "No es posible retornar el templte", http.StatusInternalServerError)
		hadlerError(w, http.StatusInternalServerError)
	}
}

func main() {

	// archivos estaticos
	statisFiles := http.FileServer(http.Dir("static"))

	port := ":3000"

	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	// mux de statuc file
	mux.Handle("/static/", http.StripPrefix("/static/", statisFiles))

	// crear servidor
	fmt.Println("El servidor esta corriendo en el puerto ", port)
	server := &http.Server{
		Addr:    "localhost" + port,
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error when start the server " + port)
	}
}
