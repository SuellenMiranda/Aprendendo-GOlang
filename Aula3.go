package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Joguinho gerando número aleatório

	// // for normal
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// // for while
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	seed := time.Now().Unix()
	rand.Seed(seed)

	numero := rand.Intn(100) + 1
	fmt.Println("Pensei em um número, de 1 a 100.")
	fmt.Println("Consegue advinhar?")
	// fmt.Println(numero)

	reader := bufio.NewReader(os.Stdin)

	sucesso := false
	for i := 0; i < 10; i++ {
		fmt.Println("Você tem: ", 10-i, "chances sobrando.")
		fmt.Println("Qual seu chute?")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		chute, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if chute == numero {
			fmt.Println("Acertouuuu")
			sucesso = true
			break
		} else if chute < numero {
			fmt.Println("Opa, abaixo do que pensei")
			continue
		} else {
			fmt.Println("Eita, essa foi acima")
		}
	}

  if !sucesso{
    fmt.Println("Você perdeu! O número era: ", numero)
  }

}
