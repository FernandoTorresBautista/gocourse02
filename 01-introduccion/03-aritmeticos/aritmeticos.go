package main

import "fmt"

func main() {
	a := 20
	b := 10

	// suma / resta
	result := a + b
	fmt.Println("Suma: ", result)
	result = a - b
	fmt.Println("Resta: ", result)
	// multiplicacion / division
	result = a * b
	fmt.Println("Multi: ", result)
	result = a / b
	fmt.Println("Division: ", result)
	var div float64 = 3.0 / 2.0
	fmt.Println("Division: ", div)

	// modulo
	result = a % b
	fmt.Println("Modulo: ", result)
}
