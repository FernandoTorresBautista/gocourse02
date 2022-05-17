package main

import "fmt"

func main() {
	/* Example: Descuento en restaurantes
	Desc de 10% hasta 100 de consumo
	Desc de 20% hasta 200 de consumo
	Desc de 30% mayor 200 de consumo
	igv 19%
	*/
	var consumo, descuento float64
	var datosDescuento string

	// input
	fmt.Print("Ingrese consumo: ")
	fmt.Scanln(&consumo)

	if consumo > 0 {
		if consumo <= 100 {
			datosDescuento = "10%"
			descuento = consumo * 0.10
		} else if consumo > 100 && consumo <= 200 {
			datosDescuento = "20%"
			descuento = consumo * 0.20
		} else if consumo > 200 {
			datosDescuento = "30%"
			descuento = consumo * 0.30
		}
	} else {
		fmt.Println("Error al ingresar el descuento")
	}

	montoDescuento := consumo - descuento
	igv := montoDescuento * 0.19
	totalPagar := montoDescuento + igv

	fmt.Println("-----------------------------------------")
	fmt.Println("Descuento que aplica: ", datosDescuento)
	fmt.Println("consumo: ", consumo)
	fmt.Println("descuento: ", descuento)
	fmt.Println("monto con descuento: ", montoDescuento)
	fmt.Println("IGV: ", igv)
	fmt.Println("total a pagar: ", totalPagar)

}
