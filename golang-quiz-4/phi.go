package main

import "fmt"

func main() {
	var n int

	var sum float64 = 0.0

	var sign float64 = 1.0

	var denominator float64 = 1.0

	var pi float64

	var i int
	// ---------------------------------------

	fmt.Scan(&n)

	for i = 0; i < n; i++ {

		sum = sum + (sign * (1.0 / denominator))

		sign = -sign

		denominator = denominator + 2.0
	}

	pi = 4.0 * sum

	fmt.Println(pi)
}
