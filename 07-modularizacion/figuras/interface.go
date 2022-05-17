package figuras

import "fmt"

type Geometria interface {
	area() string
	perimetro() string
}

func Medidas(f Geometria) {
	fmt.Println(f)
	fmt.Println("Area:", f.area())
	fmt.Println("Perimetro:", f.perimetro())
}
