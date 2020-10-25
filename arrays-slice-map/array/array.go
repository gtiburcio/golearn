package main

import "fmt"

func main() {
	var notas [3]float64
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	sum := 0.0
	for i := 0; i < len(notas); i++ {
		sum += notas[i]
	}

	media := sum / float64(len(notas))

	fmt.Printf("A média eh %.2f", media)
}
