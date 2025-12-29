package main

import (
	"fmt"
)

func main() {
	var input string
	var found bool = false

	fmt.Scan(&input)

	for _, char := range input {
		if char == 'k' || char == 'q' {
			found = true
			break
		}
	}

	fmt.Println(found)
}
