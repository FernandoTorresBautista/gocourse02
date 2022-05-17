package mensajes

import "fmt"

const mensaje = "Hola desde constante"

// func hola could be a private function

// Hola function public
func Hola() {
	fmt.Println("Hola desde el paquete mensaje")
}

func Imprimir() {
	fmt.Println(mensaje)
}

func funcionPrivada() {
	fmt.Println("Funcion privada")
}
