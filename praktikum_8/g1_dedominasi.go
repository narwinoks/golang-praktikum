package main

import "fmt"

func main() {
	var m, x, y, z int
	fmt.Scan(&m, &x, &y, &z)
	denoms := []int{x, y, z}
	for _, d := range denoms {
		count := m / d
		m = m % d
		switch {
		case count > 0:
			fmt.Printf("%d lembar uang IDR %d\n", count, d)
		}
	}
}
