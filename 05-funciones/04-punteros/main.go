package main

import "fmt"

func modificarValor(c string) {
	fmt.Printf("%p\n", &c)
	c = "Hola desde la funcion"
}

func modificarValorRefMem(c *string) {
	fmt.Printf("%p\n", &c)
	*c = "Hola desde la funcion"
}

func main() {
	c := "Hola mundo"
	fmt.Printf("%p\n", &c)
	fmt.Println("Antes de la funcion: ", c)
	modificarValor(c) // copia
	fmt.Println("Despues de la funcion: ", c)
	fmt.Printf("%p\n", &c)
	//
	fmt.Println("")
	//
	fmt.Printf("%p\n", &c)
	fmt.Println("Antes de la funcion: ", c)
	modificarValorRefMem(&c) // copia
	fmt.Println("Despues de la funcion: ", c)
	fmt.Printf("%p\n", &c)

}
