package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255 // byte é um alias uint8
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... uint8 uint16 uint32 uint64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo de int é", i1)
	fmt.Println("o tipo de i1 é", reflect.TypeOf((i1)))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("o tipo literal de 49.99 é", reflect.TypeOf(49.99)) //float64

	// boolean
	bo := true
	fmt.Println("O tipo da variável bo é", reflect.TypeOf(bo))
	fmt.Println(!bo) // false

	// string
	s1 := "Olá mundo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// string com múltiplas linhas
	s2 := `
	olá
	mundo
	!
	`
	fmt.Println("O tamanho da string é", len(s2))

	// char?
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char)) //int32
}
