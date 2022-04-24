package main

import (
	"fmt"
)

func main() {
	qtd := 6
  double(&qtd)
  fmt.Printf("%d \n", qtd)
}

func double(num *int){
  *num *=2
}
