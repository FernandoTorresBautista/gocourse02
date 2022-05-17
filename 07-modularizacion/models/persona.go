package models

type Persona struct {
	nombre string
	edad   int
}

func (p *Persona) Constructor(n string, e int) {
	p.nombre = n
	p.edad = e
}

func (p *Persona) SetNombre(n string) {
	p.nombre = n
}

func (p *Persona) SetEdad(e int) {
	p.edad = e
}

func (p *Persona) GetNombre() string {
	return p.nombre
}

func (p *Persona) GetEdad() int {
	return p.edad
}
