package main

import (
	"fmt"
)

func permutationEquation(p []int32) []int32 {
	var a []int32

	for i := 1; i <= len(p); i++ {
		x := SliceIndex(p, int32(i)) + 1
		v := SliceIndex(p, int32(x)) + 1
		a = append(a, int32(v))
	}

	return a
}

func SliceIndex(p []int32, element int32) int {
	for i, v := range p {
		if v == element {
			return i
		}
	}
	return -1
}

func main() {
	total := permutationEquation([]int32{4, 3, 5, 1, 2})
	fmt.Println(total)
}
