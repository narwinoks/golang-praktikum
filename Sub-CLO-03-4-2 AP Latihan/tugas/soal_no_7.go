package main

import (
	"fmt"
)

func main() {
	var tidakHujan, moodBagus bool

	// Membaca 2 input boolean
	// Input contoh: true true
	fmt.Scan(&tidakHujan, &moodBagus)

	// Cek kondisi: Jika tidak hujan (true) DAN mood bagus (true)
	if tidakHujan && moodBagus {
		fmt.Println("keluar jalan-jalan")
	} else {
		fmt.Println("diam di rumah saja")
	}
}
