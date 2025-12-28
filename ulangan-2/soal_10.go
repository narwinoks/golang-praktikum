package main

import "fmt"

func main() {
	var x int
	fmt.Scanln(&x)

	// Ini adalah "while" loop, akan berjalan selama x > 0
	for x > 0 {
		var digit int = x % 10

		x = x / 10

		fmt.Print(digit)
		if x > 0 {
			fmt.Print(" ")
		}
	}

	fmt.Println()
}
