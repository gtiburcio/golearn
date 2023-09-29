package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		fmt.Println("You entered:", text)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
