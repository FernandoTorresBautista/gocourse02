package main

import "fmt"

func main() {
	// not
	fmt.Println(!true)
	// and
	fmt.Println(true && false)
	fmt.Println(true && true)
	// or
	fmt.Println(true || false)
	fmt.Println(true || true)
	//
	b := 2
	r := b == 2 || (b < 5)
	fmt.Println(r)
}
