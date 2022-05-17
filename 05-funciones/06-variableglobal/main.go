package main

import "fmt"

var mensaje string

func f1() {
	mensaje = "desde f1"
	fmt.Println(mensaje)
}

func f2() {
	mensaje = "desde f2"
	fmt.Println(mensaje)
}

func main() {
	mensaje = "desde main"
	fmt.Println(mensaje)

	f1()
	f2()

	fmt.Println(mensaje)
}
