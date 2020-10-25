package main

import "time"

func main() {
	t := time.Now()
	switch { // switch true
	case t.Hour() < 12:
		println("Bom dia!")
	case t.Hour() < 18:
		println("Boa tarde!")
	default:
		println("Boa noite!")
	}
}
