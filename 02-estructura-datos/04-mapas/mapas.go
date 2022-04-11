package main

import "fmt"

func main() {
	dias := make(map[int]string)

	fmt.Println(dias)

	// agregar datos
	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miercoles"
	dias[5] = "Jueves"
	dias[6] = "Viernes"
	dias[7] = "Sabado"
	fmt.Println(dias)
	dias[7] = "SABADO"
	fmt.Println(dias)
	delete(dias, 1) // delete domingo
	fmt.Println(dias, len(dias))

	// nuevo mapa
	estudiantes := make(map[string][]int)
	estudiantes["Fer"] = []int{13, 15, 16}
	estudiantes["Fer2"] = []int{14, 13, 17}
	fmt.Println(estudiantes)
	fmt.Println(estudiantes["Fer"])
	fmt.Println(estudiantes["Fer"][1])
}
