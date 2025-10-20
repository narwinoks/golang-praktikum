package main

import "fmt"

func menghitung(jumlahPesanan []int) int {
	//PRODUC PRICE ADA DISOAL
	harga := []int{15000, 25000, 5000, 10000}

	totalHarga := 0

	for i, jumlah := range jumlahPesanan {
		totalHarga += jumlah * harga[i]
	}

	return totalHarga
}
func main() {
	var pesanan1 []int = []int{5, 0, 0, 0}
	total1 := menghitung(pesanan1)

	var pesanan2 []int = []int{1, 2, 3, 4}
	var total2 int = menghitung(pesanan2)

	var pesanan3 []int = []int{1, 1, 1, 1}
	var total3 int = menghitung(pesanan3)

	var pesanan4 []int = []int{15, 32, 18, 100}
	var total4 int = menghitung(pesanan4)

	fmt.Println("PESANAN PERTAMA :", "IDR ", total1)
	fmt.Println("PESANAN KEDUA :", "IDR ", total2)
	fmt.Println("PESANAN KETIGA :", "IDR ", total3)
	fmt.Println("PESANAN KEEMPAT :", "IDR ", total4)

}
