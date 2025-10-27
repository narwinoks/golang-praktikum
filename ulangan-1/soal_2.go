package main

import "fmt"

func hitungKeliling(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}

func main() {
	var panjang, lebar, hasil int
	fmt.Scanln(&panjang)
	fmt.Scanln(&lebar)

	hasil = hitungKeliling(panjang, lebar)

	fmt.Println(hasil)
}
