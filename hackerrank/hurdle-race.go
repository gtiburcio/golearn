package main

func main() {
	array := [5]int32{1, 6, 3, 5, 2}
	hurdleRace(4, array[:])
}

func hurdleRace(k int32, height []int32) int32 {
	max := 0
	i int32
	for i = 0; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
	}
	println(max)
	return 0
}
