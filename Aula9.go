package main

import (
	"fmt"
  "time"
)

func simples(msg string, c chan string){
  for i := 0 ; i < 5 ; i++{
    c <- fmt.Sprint(msg)
  }
}
func main() {
  c := make(chan string)
  go simples("Ping", c)
  for i := 0 ; i < 5 ; i++{
    fmt.Printf("Received: %q \n", <- c)
    fmt.Printf("Pong\n")
    time.Sleep(1 * time.Second)
    
  }
  time.Sleep(10 * time.Second)
  fmt.Println("Fim do programa")
}
