package main

func reverseArray(a []int32) []int32 {

	var slice []int32

	for i := len(a) - 1; i >= 0; i-- {
		slice = append(slice, a[i])
	}

	return slice
}

func main() {

}
