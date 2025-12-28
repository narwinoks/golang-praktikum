package main

import (
	"fmt"
)

func main() {
	var m float64
	var c float64

	fmt.Scanln(&m)
	fmt.Scanln(&c)
	const epsilon = 1e-9

	if abs(c) < epsilon {
		fmt.Println("melewati")
	} else {
		fmt.Println("tidak melewati")
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
