package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var currentGPA float64
	var maxGPA float64 = 0.0
	var count int = 0

	i := 0
	for i < n {
		fmt.Scan(&currentGPA)

		if currentGPA > maxGPA {
			maxGPA = currentGPA
			count = 1
		} else if currentGPA == maxGPA {
			count++
		}

		i++
	}

	fmt.Printf("%d %.1f\n", count, maxGPA)
}
