package main
import (
	"fmt"
  "io/ioutil"
  "net/http"
)
func size(url string, c chan int) {
	fmt.Printf("Pegando o site:%s\n", url)
  response, _ := http.Get(url)
  defer response.Body.Close()
  body, _ := ioutil. ReadAll(response.Body)
  fmt.Printf("Tamanho=%d da pag%s\n", len(body), url)
  c <- len(body)
}
func main(){
  c := make(chan int)
  go size("https://uvv.br/", c)
  go size("https://ufes.br/", c)
  go size("https://ifes.edu.br/", c)
  var total int
  total += <- c
  total += <- c
  total += <- c
  fmt.Printf("Tamanho Total = %d\n", total)
}
