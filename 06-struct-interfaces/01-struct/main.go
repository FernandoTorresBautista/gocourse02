package main

import "fmt"

type Persona struct {
	Name string
	Age  int
}

func (p *Persona) imprimir() {
	fmt.Println("Name:", p.Name, "Age:", p.Age)
}

func (p *Persona) setName(newName string) {
	p.Name = newName
}

// herencia
type Empleado struct {
	Persona
	Sueldo float64
}

func main() {
	p1 := Persona{}
	p1.imprimir()

	p2 := Persona{"Fer", 26}
	p2.imprimir()

	// p2.Name = "Fer2"
	p2.setName("Fer2")
	p2.imprimir()

	p3 := Persona{
		Name: "Fer3",
		Age:  26,
	}
	p3.imprimir()

	// herencia
	emp1 := Empleado{
		Sueldo: 500,
	}
	emp1.Name = "Fer4"
	emp1.Age = 16
	fmt.Println(emp1)
	emp1.imprimir()
}
