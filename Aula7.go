//de nada

package main

import (
	"fmt"
  "main/data"
)

type evento struct{
  title string
  data.Date
}

func main() {  
  // date := data.Date{}
  // date.SetDay(20)
  // err := date.SetMonth(12)
  // if err != nil{
  //   log.Fatal(err)
  // }
  // date.SetYear(2023)

  formatura := evento{title: "Dia da Formatura:"}
  formatura.SetDay(11)
  formatura.SetMonth(02)
  formatura.SetYear(2024)
  
	fmt.Printf("%v \n",formatura)
}
