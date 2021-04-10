package main

import (
	"fmt"
)

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s, %s\n", p1, p2)
}

func main() {
	f1()
	f2("eu sou", "a função 2")
}
