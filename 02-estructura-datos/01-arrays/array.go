package main

import "fmt"

func main() {
	// array
	var numeros [5]int
	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	fmt.Println(numeros)
	fmt.Println(numeros[1])

	// array con valores
	nombres := [3]string{"uno", "dos", "tres"}
	fmt.Println(nombres)

	// array sin longitud
	colores := [...]string{"Rojo", "Verde", "Negro", "Azul"}
	fmt.Println(colores, len(colores))

	// indices definidos
	monedas := [...]string{
		0: "dollar",
		2: "pesos",
		5: "euro",
	}
	fmt.Println(monedas, len(monedas))
	fmt.Println(monedas[0:3])
	fmt.Println(monedas[3:])
}
