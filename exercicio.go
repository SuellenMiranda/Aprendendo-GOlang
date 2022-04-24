package main

import (
	"fmt"
)

type aluno struct{
  nome string
  matricula int
  endereco
}

type endereco struct{
  numero int
  rua string
}

var m int = 1
func print(a aluno) {
  fmt.Printf("M: %d, S: %s \n", a.matricula, a.nome)
  fmt.Printf("Endereço: %s\n", a.rua)
}

func novoAluno(nome string)aluno {
  a := aluno{matricula:m, nome: nome, }
  a.endereco = endereco{numero: 123, rua: "Rodosol"}
  m++
  return a
}

func main() {
  a:= novoAluno("aFulano")
  print (a)
}
