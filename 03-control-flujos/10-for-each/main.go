package main

import "fmt"

func main() {

	nombres := [...]string{"Fer", "Fer2", "Fer3"}
	// for i := 0; i < len(nombres); i++ {
	// 	fmt.Println(nombres[i])
	// }

	// for i, v := range nombres {
	for _, v := range nombres {
		// fmt.Println(i, v)
		fmt.Println(v)
	}
}
