package main

import (
	"fmt"
)

func main() {
	a := []int32{1, 2, 2, 3, 1, 2}
	pickingNumbers(a)
}

func pickingNumbers(a []int32) int32 {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	return 0
}
