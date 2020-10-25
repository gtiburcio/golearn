package main

import "time"

func main() {

	i := 1
	for i <= 10 {
		println(i)
		i++
	}

	println("-----------------------------")

	for i := 0; i <= 20; i += 2 {
		println(i)
	}

	println("-----------------------------")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			print("Par ")
		} else {
			print("Impar ")
		}
	}

	println("\n-----------------------------")

	for {
		println("Para sempre...")
		time.Sleep(time.Second)
	}
}
