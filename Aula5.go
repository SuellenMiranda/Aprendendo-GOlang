package main

import (
	"fmt"
)

func main() {
  pesos := make([]float64, 0)

  var peso float64
  for peso >= 0 {
    fmt.Printf("Digite um peso:\n")
    fmt.Scanf("%f", &peso)
    
    pesos = append(pesos, peso)
  }

  subSlice := pesos [ 0 : len(pesos) -1]
  var soma float64
  for _,p := range subSlice{
    soma += p
  }
  
  fmt.Printf("%f \n", subSlice)
  fmt.Printf("Media: %f \n", soma/ float64(len(subSlice)))
}
