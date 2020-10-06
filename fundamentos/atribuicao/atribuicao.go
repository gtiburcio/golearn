package main

func main() {
	var b byte = 3
	println(b)

	i := 3
	i += 3
	i -= 3
	i /= 2
	i *= 2
	i %= 2
	println(i)

	x, y := 1, 2
	println(x, y)

	x, y = y, x
	println(x, y)
}
