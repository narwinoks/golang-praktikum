package main

import "fmt"

func hitungLuas(panjang float64, lebar float64) float64 {
	return panjang * lebar
}
func main() {
	var panjang, lebar, luas float64
	fmt.Scanln(&panjang, &lebar)
	luas = hitungLuas(panjang, lebar)
	fmt.Println(luas)
}
