package figuras

import "fmt"

type Cuadrado struct {
	Lado float64
}

func (c *Cuadrado) area() string {
	return fmt.Sprintf("Area: %f", c.Lado*c.Lado)
}

func (c *Cuadrado) perimetro() string {
	return fmt.Sprintf("perimetro: %f", (2*c.Lado)+(2*c.Lado))
}
