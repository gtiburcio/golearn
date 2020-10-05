package main

import "fmt"

func main() {

	x := 3.141516

	// fmt.Println("O valor de x é " + x)
	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é " + xs)

	// ou

	fmt.Println("O valor de x é", xs)

	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %.1f %t %s", a, b, c, d)
}
