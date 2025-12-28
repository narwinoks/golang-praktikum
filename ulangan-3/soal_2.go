package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scanln(&a, &b, &c, &d)

	var x, y int

	x = (a * d) + (c * b)
	y = b * d
	fmt.Printf("penjumlahan: %d / %d\n", x, y)

	x = (a * d) - (c * b)
	y = b * d
	fmt.Printf("pengurangan: %d / %d\n", x, y)

	x = a * c
	y = b * d
	fmt.Printf("perkalian: %d / %d\n", x, y)

	x = a * d
	y = b * c
	fmt.Printf("pembagian: %d / %d\n", x, y)
}
