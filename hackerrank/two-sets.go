package main

import (
	"fmt"
	"math"
)

func getTotalX(a []int32, b []int32) int32 {
	numberOfValues := int32(0)

	maxA := findMaxValueInArray(a)
	minB := findMinValueInArray(b)

	for i := maxA; i <= minB; i++ {

		eligible := false
		for _, valueA := range a {
			if i%valueA != 0 {
				eligible = false
				break
			}
			eligible = true
		}

		if eligible == false {
			continue
		}

		eligible = false

		for _, valueB := range b {
			if valueB%i != 0 {
				eligible = false
				break
			}
			eligible = true
		}

		if eligible == false {
			continue
		}

		numberOfValues++

	}

	return numberOfValues
}

func findMaxValueInArray(a []int32) int32 {
	max := int32(0)

	for _, value := range a {
		if value >= max {
			max = value
		}
	}
	return max
}

func findMinValueInArray(a []int32) int32 {
	min := int32(math.MaxInt32)

	for _, value := range a {
		if value <= min {
			min = value
		}
	}
	return min
}

func main() {
	total := getTotalX([]int32{2, 6}, []int32{24, 36})
	fmt.Println(total)
}
