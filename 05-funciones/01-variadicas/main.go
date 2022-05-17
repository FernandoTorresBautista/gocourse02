package main

import "fmt"

func sumar(nombre string, numeros ...int) (name string, result int) {
	name = nombre
	for _, v := range numeros {
		result += v
	}
	return
}

func main() {
	message, result := sumar("fer")
	fmt.Println(message, result)
	fmt.Println(sumar("fer", 1, 2, 3, 4, 5))
}
