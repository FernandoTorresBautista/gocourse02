package main

import "fmt"

func main() {
	var name string
	name = "any name"

	var name2 string = "any name"

	edad := 20

	fmt.Println(name, name2, edad)

	var a string
	var b float64
	var c string
	var d bool
	fmt.Println(a, "-", b, "-", c, "-", d)

	const pi = 3.1416
	fmt.Println(pi)
}
