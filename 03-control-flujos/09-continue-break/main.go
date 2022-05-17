package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 5 {
			fmt.Println("Salta la iteracion")
			continue
		}
		if i == 8 {
			fmt.Println("romper bucle iteracion 8")
			break
		}
		fmt.Println(i)
	}
}
