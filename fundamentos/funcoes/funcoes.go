package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}

func doisRetornos(valor1 int, valor2 int) (int, int) {
	return somar(valor1, valor2), valor1 * valor2
}
