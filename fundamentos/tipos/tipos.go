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

	// somente positivos... uint8 uint16 uint32 uint64
	const b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// positivos e negativos... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))                 // float32
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64
	fmt.Println("O máximo do float64 é", math.MaxFloat64)

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println("O valor de bo é", bo)
	fmt.Println(!bo)

	// string
	s1 := "Olá meu nome é Tiburcio"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// string com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Tiburcio`

	println("O tamanho da string agora é", len(s2))
}
