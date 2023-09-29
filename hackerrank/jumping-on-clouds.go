package main

import "fmt"

func jumpingOnClouds(c []int32, k int32) int32 {

	var total int32 = 100

	returnBegining := false

	var i int32 = 0

	var size int32 = int32(len(c))

	f := 0

	for returnBegining != true {

		f++

		i = (i + k) % size

		fmt.Printf("turn %d value %d - %d\n", f, i, c[i])

		value := c[i]

		if value == 1 {
			total -= 2 + 1
		} else {
			total--
		}

		if i == 0 {
			returnBegining = true
		}

	}

	return total
}

func main() {
	// total := jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 1, 0}, 2)
	// fmt.Println(total)

	total2 := jumpingOnClouds([]int32{1, 1, 1, 0, 1, 1, 0, 0, 0, 0}, 3)
	fmt.Println(total2)

	// total3 := jumpingOnClouds([]int32{0, 0, 1, 0}, 2)
	// fmt.Println(total3)
}
