package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.43,
		"Gabriela Silva": 2332.34,
		"Pedro Junio":    2313.32,
	}

	funcsESalarios["Rafael Filho"] = 1350.0

	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
