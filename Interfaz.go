package main

import "fmt"

/* Declaración de una interfaz tipo FiguraGeométrica*/
type FiguraGeometrica interface {
	area() float64
	//Aquí podrían ir más métodos
}

//Usuario ...
type Usuario interface {
	Nombre() string
	Email() string
}

//Persona ...
type Persona struct {
	nombre string
	email  string
	edad   int
}

//Jugador ...
type Jugador struct {
	nombre string
	email  string
	edad   int
}

//Nombre ...
func (p Persona) Nombre() string {
	return p.nombre
}

//Email ...
func (p Persona) Email() string {
	return p.email
}

//Nombre ...
func (p Jugador) Nombre() string {
	return p.nombre
}

//Email ...
func (p Jugador) Email() string {
	return p.email
}

//Presentarse ...
func Presentarse(p Usuario) {
	fmt.Println("Nombre:", p.Nombre())
	fmt.Println("Email:", p.Email())
}

//Rectangulo ...
type Rectangulo struct {
	base, altura float64
}

/* Declaración de registro tipo Trapecio */
type Trapecio struct {
	base_mayor, base_menor, altura float64
}

/* Método para los registros Rectangulo implementando la interfaz */
func (figura Rectangulo) area() float64 {
	return figura.base * figura.altura
}

/* Método para los registros Trapecio implementando la interfaz */
func (figura Trapecio) area() float64 {
	return (figura.base_mayor + figura.base_menor) * figura.altura / 2
}

/* Definición de un método para Figura */
func dameArea(fig FiguraGeometrica) float64 {
	return fig.area()
}
