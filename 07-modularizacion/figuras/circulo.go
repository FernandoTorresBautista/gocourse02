package figuras

import "fmt"

type Circulo struct {
	Radio float64
}

func (c *Circulo) area() string {
	return fmt.Sprintf("Area: %f", 3.1416*(c.Radio*c.Radio))
}

func (c *Circulo) perimetro() string {
	return fmt.Sprintf("perimetro: %f", (2*3.1416)*c.Radio)
}
