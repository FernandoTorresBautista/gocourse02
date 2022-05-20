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
var temp = template.Must(template.New("T").ParseFiles("templates/index.html", "templates/base.html"))
var temp2 = template.Must(template.New("T").ParseGlob("newtemp/**/*.html"))
var errorTemplate = template.Must(template.ParseFiles("newtemp/error/error.html"))

// handlers

func Index(w http.ResponseWriter, r *http.Request) {
	c1 := Curso{"Go"}
	c2 := Curso{"Java"}
	c3 := Curso{"Python"}
	c4 := Curso{"JS"}
	// fmt.Fprintln(w, "Hola mundo")
	// temp, err := template.ParseFiles("templates/index.html")
	// temp, err := template.ParseFiles("templates/index.html", "templates/base.html")
	// temp, err := template.New("index.html").ParseFiles("templates/index.html", "templates/base.html")
	// temp := template.Must(template.New("index.html").ParseFiles("templates/index.html", "templates/base.html"))

	cursos := []Curso{c1, c2, c3, c4}
	user := Usuarios{"Fer", 26, true, false, cursos}

	// usar el temp con todas las variables
	if err := temp.ExecuteTemplate(w, "index.html", user); err != nil {
		log.Fatal(err)
	}
}

func Funciones(w http.ResponseWriter, r *http.Request) {

	funciones := template.FuncMap{
		"saludar": Saludar,
	}
	temp2, err := template.New("index2.html").Funcs(funciones).ParseFiles("templates/index2.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
	} else {
		err := temp2.Execute(w, nil)
		if err != nil {
			log.Fatal("Error ", err)
		}
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
	if err := temp2.ExecuteTemplate(w, name, data); err != nil {
		// http.Error(w, "No es posible retornar el templte", http.StatusInternalServerError)
		hadlerError(w, http.StatusInternalServerError)
	}
}

func main() {
	port := ":3000"

	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/funciones", Funciones)
	mux.HandleFunc("/registro", Registro)

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
