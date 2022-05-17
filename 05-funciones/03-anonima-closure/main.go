package main

import (
	"fmt"
	"strings"
)

func repeat(n int) func(c string) string {
	// la funcion retorna otra funcion
	return func(c string) string {
		return strings.Repeat(c, n)
	}
}

func main() {

	// funciones anonimas
	func() {
		fmt.Println("Hola desde la funcion anonima(closure)")
	}()

	myfunc := func() {
		fmt.Println("Hola desde la funcion anonima(closure)")
	}
	myfunc()

	myfunc2 := func(name string) string {
		return fmt.Sprintf("Hola, %s desde la funcion anonima(closure)", name)
	}
	fmt.Println(myfunc2("fer"))

	// funciones closure
	repeat3 := repeat(3)
	fmt.Println(repeat3("fer"))
	repeat5 := repeat(5)
	fmt.Println(repeat5("fer"))
}
