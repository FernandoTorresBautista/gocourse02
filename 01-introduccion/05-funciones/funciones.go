package main

import "fmt"

func saludar(nombre string) {
	fmt.Println("Hola", nombre)
}

func sumar(a, b int) int {
	return a + b
}

func main() {
	saludar("fer")

	fmt.Println(sumar(2, 3))
}
