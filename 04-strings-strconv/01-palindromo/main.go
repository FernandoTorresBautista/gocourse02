package main

import (
	"fmt"
	"strings"
)

func reverse(c string) string {
	arr := strings.Split(c, "")
	arrinv := make([]string, 0)
	// fmt.Println(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		arrinv = append(arrinv, arr[i])
	}
	// fmt.Println(arrinv)
	return strings.Join(arrinv, "")
}

func esPalindromo(p string) bool {
	// fmt.Println(p)
	palabra := strings.ToLower(p)
	// fmt.Println(palabra)
	// palabra = strings.Replace(palabra, "z", "s", -1)
	palabra = strings.ReplaceAll(palabra, " ", "")
	palabraInvertida := reverse(palabra)
	// fmt.Println(palabraInvertida)
	return palabra == palabraInvertida
}

// palindromo - palabras que leen igual al derecho y al reves
func main() {
	if esPalindromo("Luz azul") {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}
