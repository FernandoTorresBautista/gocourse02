package main

import (
	"fmt"
	"log"
	"net/http"
)

// handlers
func Hola(w http.ResponseWriter, r *http.Request) {
	fmt.Println("El metodo es ", r.Method)
	fmt.Fprintln(w, "Hola mundo")
}

func PageNF(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func Error(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Mensaje de error", http.StatusNotFound)
}

func Saludar(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	q := r.URL.RawQuery
	mapq := r.URL.Query()
	fmt.Println(q)
	fmt.Println(mapq)
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	fmt.Fprintf(w, "Hola %s, tu edad es %s", name, age)
}

func main() {

	port := ":3000"
	// router
	// http.HandleFunc("/", Hola)
	// http.HandleFunc("/page", PageNF)
	// http.HandleFunc("/error", Error)
	// http.HandleFunc("/saludar", Saludar)

	// mux - ruta asociada a un handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/page", PageNF)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/saludar", Saludar)

	// crear servidor
	fmt.Println("El servidor esta corriendo en el puerto ", port)
	// if err := http.ListenAndServe("localhost"+port, nil); err != nil { // usamos nil para mux por defecto
	// if err := http.ListenAndServe("localhost"+port, mux); err != nil {
	// 	log.Fatal("Error when start the server " + port)
	// }
	server := &http.Server{
		Addr:    "localhost" + port,
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error when start the server " + port)
	}
}
