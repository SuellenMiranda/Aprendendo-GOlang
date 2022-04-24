package main

import (
	"fmt"
  "math"
)

func main() {
  c := circulo{raio:4.3}
  r := rect{base:5.2, altura:12}

  medidas(c)
  medidas(r)
	fmt.Println("Hello! Word!")
}

func medidas(fg formaGeometrica){
	fmt.Printf("Area=%f, Perimetro=%f\n", fg.area(), fg.perimetro())
}

type formaGeometrica interface{
  area() float64
  perimetro() float64
}

//Tipo Circulo que  implementa a Interface
type circulo struct{
  raio float64
}

func (c circulo) area() float64{
  return math.Pi * c.raio * c.raio
}
func (c circulo) perimetro() float64{
  return 2 * math.Pi * c.raio
}

//Tipo Retangulo que implementa a Interface

type rect struct{
  base float64
  altura float64
}

func (r rect) area() float64{
  return math.Pi * r.altura * r.base
}
func (r rect) perimetro() float64{
  return 2*r.base + 2*r.altura
}
