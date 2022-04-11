package main

import "fmt"

func main() {
	// slice
	numeros := []int{1, 2, 3}
	fmt.Println(numeros, len(numeros))

	numeros = append(numeros, 4)
	fmt.Println(numeros, len(numeros))

	// sub slice
	subSlice := numeros[:2]
	numeros[0] = 100
	fmt.Println(subSlice)
	fmt.Println(numeros)

	// punteros
	// longitud
	//capacidad

	meses := []string{"Enero", "Febrero", "Marzo"}
	fmt.Printf("Len %v, Cap %v, %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "Abril")
	fmt.Printf("Len %v, Cap %v, %p \n", len(meses), cap(meses), meses)
}
