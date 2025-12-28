package main

import "fmt"

func main() {
	var totalKembalian int

	fmt.Scanln(&totalKembalian)

	var sepuluhRibu int = totalKembalian / 10000

	var sisa int = totalKembalian % 10000

	var limaRibu int = sisa / 5000

	sisa = sisa % 5000

	var seribu int = sisa / 1000

	fmt.Println(sepuluhRibu, limaRibu, seribu)
}
