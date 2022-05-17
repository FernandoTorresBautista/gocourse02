package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Ups!, al parecer el programa no finalizo de forma correcta")
		}
	}()

	if file, err := os.Open("./hola_.txt"); err != nil {
		// fmt.Println(err)
		// return
		panic("no fue posible obtener el archivo")
	} else {
		defer func() {
			fmt.Println("Cerrando el archivo!")
			file.Close()
		}()
		content := make([]byte, 254)
		long, _ := file.Read(content)
		contenidoArchivo := string(content[:long])
		fmt.Println(contenidoArchivo)
	}
	// file.Close()

}
