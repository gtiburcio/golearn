package main

func main() {
	number := 8

	atual, ant, antAnt := 0, 0, 1
	for number > 0 {
		atual = ant + antAnt
		antAnt = ant
		ant = atual
		println(atual)
		number--
	}
}

// recursivo
func fibo(n int) int {
	if n < 2 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}
