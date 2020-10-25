package main

import "fmt"

func main() {
	numeros := [...]int{2, 4, 6, 8, 10} // compilador conta

	for i, numero := range numeros {
		fmt.Printf("%d and %d\n", i, numero)
	}

	for _, numero := range numeros {
		fmt.Printf("%d\n", numero)
	}
}
