package main

func notaParaConceito(n float64) string {

	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	n1 := notaParaConceito(10)
	n2 := notaParaConceito(8)
	n3 := notaParaConceito(6)
	n4 := notaParaConceito(3)

	println(n1)
	println(n2)
	println(n3)
	println(n4)
}
