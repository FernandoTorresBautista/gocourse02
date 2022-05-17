package main

import "fmt"

func main() {
	// bucle infinito
	// for {
	// 	fmt.Println("Hola mundo")
	// }

	numeros := 12455
	c := 0
	// bucle tipo while
	for numeros > 0 {
		numeros /= 10
		c++
	}
	fmt.Println("Cantidad de digitos: ", c)

	// bucle for
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}

	// for c, x := range aux[interface]interface
}
