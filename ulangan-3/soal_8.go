package main

import "fmt"

func main() {
	var tabungan, total int
	var jumlah int
	fmt.Scanln(&tabungan)
	fmt.Scanln(&jumlah)
	for i := 1; i <= jumlah; i++ {
		total = total + tabungan
		tabungan += 2500
	}
	fmt.Println(total)
}
