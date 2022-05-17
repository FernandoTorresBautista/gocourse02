package main

import "fmt"

// reto

type Figura interface {
	area() string
	perimetro() string
}
type Cuadrado struct {
	ancho float64
	alto  float64
}

type Circulo struct {
	radio float64
}

func (c *Cuadrado) area() string {
	return fmt.Sprintf("Area: %f", c.ancho*c.alto)
}

func (c *Circulo) area() string {
	return fmt.Sprintf("Area: %f", 3.1416*(c.radio*c.radio))
}

func (c *Cuadrado) perimetro() string {
	return fmt.Sprintf("perimetro: %f", (2*c.ancho)+(2*c.alto))
}

func (c *Circulo) perimetro() string {
	return fmt.Sprintf("perimetro: %f", (2*3.1416)*c.radio)
}

//
type Animal interface {
	mover() string
}

type Perro struct{}
type Pez struct{}
type Pajaro struct{}

func (*Perro) mover() string {
	return "Soy perro y camino"
}

func (*Pez) mover() string {
	return "Soy pez y nado"
}

func (*Pajaro) mover() string {
	return "Soy pajaro y vuelo"
}

// acciones de interfaces
func moverAnimal(animal Animal) {
	fmt.Println(animal.mover())
}

func area(figura Figura) {
	fmt.Println(figura.area())
}

func perimetro(figura Figura) {
	fmt.Println(figura.perimetro())
}

func main() {
	d1 := Perro{}
	moverAnimal(&d1)
	d2 := Pez{}
	moverAnimal(&d2)
	d3 := Pajaro{}
	moverAnimal(&d3)

	//
	f1 := Cuadrado{10, 10}
	f2 := Circulo{2}
	fmt.Println("Cuadrado")
	area(&f1)
	perimetro(&f1)
	fmt.Println("Circulo")
	area(&f2)
	perimetro(&f2)
}
