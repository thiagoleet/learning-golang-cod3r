package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	// bloco de inicialização
	if i := numeroAleatorio(); i > 5 { // também suportado no switch
		fmt.Println("Ganhou!")
	} else {
		fmt.Println("Perdeu!")
	}

	// fmt.Println(i) ERRO i está indefinido
}