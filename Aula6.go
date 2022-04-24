package main

import (
	"fmt"
)

type aluno struct{
  matricula int
  senha string
  status bool
  endereco
}

type endereco struct{
  numero int
  rua string
}

var m int = 1
func print(a aluno) {
  fmt.Printf("M: %d, S: %t \n", a.matricula, a.status)
  fmt.Printf("Endereço: %s\n", a.rua)
}

func novoAluno(senha string)aluno {
  a := aluno{matricula:m, senha: senha, status: true}
  a.endereco = endereco{numero: 123, rua: "Rodosol"}
  m++
  return a
}

func expulsaoAluno(a aluno){
  a.status = false
}

func main() {
  a:= novoAluno("123456")
  expulsaoAluno(a)
  print (a)
}
