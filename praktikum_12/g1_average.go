package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0
	count := 0
	for n > 0 {

		digit := n % 10

		if digit%2 != 0 {
			sum += digit
			count++
		}
		n = n / 10
	}
	if count > 0 {
		avg := float64(sum) / float64(count)

		fmt.Printf("%g\n", avg)
	}

	if count == 0 {
		fmt.Println("NaN")
	}
}
