package main

import (
	"fmt"
)

func main() {
	var value int
	println("digite o valor")
	fmt.Scan(&value)
	sum := 0
	for i := 1; i <= value; i++ {
		if value%i == 0 {
			sum++
		}
	}

	if sum == 2 {
		println("primo")
	} else {
		println("não primo")
	}

}
