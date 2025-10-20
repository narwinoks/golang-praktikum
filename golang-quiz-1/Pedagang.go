package main

import "fmt"

func persen(persenHargaJual int, harga int) float64 {
	return float64(persenHargaJual)/100.0*float64(harga) + float64(harga)
}
func main() {
	var persenHargaJual, hargaBarang1, hargaBarang2, hargaBarang3 int
	fmt.Print("Masukkan PersenHarga ,harga barang 1,2,3: ")
	fmt.Scan(&persenHargaJual, &hargaBarang1, &hargaBarang2, &hargaBarang3)
	fmt.Println("Harga Jual Barang 1 ", persen(persenHargaJual, hargaBarang1))
	fmt.Println("Harga Jual Barang 2 ", persen(persenHargaJual, hargaBarang2))
	fmt.Println("Harga Jual Barang 3 ", persen(persenHargaJual, hargaBarang3))
}
