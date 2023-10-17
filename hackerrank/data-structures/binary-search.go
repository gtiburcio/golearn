package main

import (
	"sort"
)

func main() {
	binarySearch([]int{1}, 1)
}

func binarySearch(ar []int, x int) int {
	sort.Ints(ar)

	min := 0
	max := len(ar) - 1

	for min <= max {
		mid := min + (max-min)/2
		if ar[mid] > x {
			max = mid - 1
		} else if ar[mid] < x {
			min = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
