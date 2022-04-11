package main

import "fmt"

func main() {
	hola := "hola"
	mundo := "mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Fer"
	edad := 10

	fmt.Printf("Hola %s y su edad es %d\n", nombre, edad)
	fmt.Printf("Hola %v y su edad es %v\n", nombre, edad)

	mensaje := fmt.Sprintf("Hola %s y su edad es %d\n", nombre, edad)
	fmt.Println(mensaje)
	// tipo de dato
	fmt.Printf("nombre: %T\n", nombre)
	fmt.Printf("edad: %T\n", edad)

	fmt.Print("Ingrese otro nombre ")
	fmt.Scanln(&nombre)

	fmt.Println("Otro nombre", nombre)
}
