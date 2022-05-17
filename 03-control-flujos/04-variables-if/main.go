package main

import "fmt"

func main() {
	if nombre, edad := "Fer", 26; nombre == "Fer" {
		fmt.Println("Hola ", nombre)
	} else {
		fmt.Printf("Nombre %s - Edad: %d\n", nombre, edad)
	}

	users := make(map[string]string)
	users["Fer1"] = "fer1@test.com"
	users["Fer2"] = "fer2@test.com"

	if correo, ok := users["Fer"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No existe")
	}
}
