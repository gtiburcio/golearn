package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 23423.43,
			"Guga Pereira":   2324.12,
		},
		"J": {
			"José João": 12434.43,
		},
		"P": {
			"Pedro Junior": 1200.00,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
