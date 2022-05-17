package main

import (
	"fmt"
	"modularizacion/figuras"
	"modularizacion/mensajes"
	"modularizacion/models"

	"github.com/donvito/hellomod"
)

// correr en la carpeta donde esta este archivo
// go mod init modularizacion
// podremos usar las subcarpetas como
//		modularizacion/[package]
func main() {
	fmt.Println("Mensajes")
	mensajes.Hola()
	mensajes.Imprimir()
	// modulariza figuras
	fmt.Println("Figuras geometricas")
	cua1 := figuras.Cuadrado{Lado: 8}
	figuras.Medidas(&cua1)
	cir1 := figuras.Circulo{Radio: 2}
	figuras.Medidas(&cir1)
	// encapsulamiento de interfaces como clases
	fmt.Println("Encapsulamiento")
	// p1 := models.Persona{}
	p1 := models.Persona{}
	p1.Constructor("Fer", 26)
	fmt.Println(p1)
	p1.SetNombre("NewFer")
	p1.SetEdad(27)
	fmt.Println(p1)

	// using package of third parts
	fmt.Println("\nSay hello")
	hellomod.SayHello()

}
