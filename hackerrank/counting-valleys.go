package main

func main() {
	println(countingValleys(8, "DDUUUUDD"))
}

func countingValleys(steps int32, path string) int32 {
	valleys := 0
	lastLevel := 0
	actualLevel := 0
	for _, r := range path {
		c := string(r)
		lastLevel = actualLevel
		if c == "D" {
			actualLevel--
		} else {
			actualLevel++
		}
		if lastLevel < 0 && actualLevel == 0 {
			valleys++
		}
	}
	return int32(valleys)
}
