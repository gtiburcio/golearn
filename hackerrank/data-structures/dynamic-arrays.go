package main

import "fmt"

func dynamicArray(n int32, queries [][]int32) []int32 {
	arr := make([][]int32, n)

	var result []int32

	var lastAnswer int32 = 0

	for i := int32(0); i < int32(len(queries)); i++ {

		option := queries[i][int32(0)]

		x := queries[i][int32(1)]

		y := queries[i][int32(2)]

		idx := (x ^ lastAnswer) % n

		if option == int32(1) {
			arr[idx] = append(arr[idx], y)
		} else {
			lastAnswer = arr[idx][y%int32(len(arr[idx]))]
			result = append(result, lastAnswer)
		}
	}

	return result
}

func main() {
	result := dynamicArray(2, [][]int32{
		{1, 0, 5},
		{1, 1, 7},
		{1, 0, 3},
		{2, 1, 0},
		{2, 1, 1},
	})
	fmt.Println(result)
}
