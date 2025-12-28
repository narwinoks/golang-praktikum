package main

import "fmt"

func main() {
	var n int
	var totalTeks int = 0
	var totalGambar int = 0

	// 1. Baca jumlah hari (n)
	fmt.Scanln(&n)

	// 2. Loop sebanyak n kali
	for i := 0; i < n; i++ {
		// Siapkan variabel untuk menampung input harian
		var teks, gambar int

		// 3. Baca 2 angka (teks dan gambar) dari baris saat ini
		fmt.Scanln(&teks, &gambar)

		// 4. Tambahkan ke total
		totalTeks = totalTeks + teks
		totalGambar = totalGambar + gambar
	}

	// 5. Cetak hasil akhirnya
	//    (Println otomatis memberi spasi di antara dua variabel)
	fmt.Println(totalTeks, totalGambar)
}
