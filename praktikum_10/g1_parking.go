package main

import "fmt"

func main() {
	var h1, m1, h2, m2 int
	fmt.Scan(&h1, &m1)
	fmt.Scan(&h2, &m2)

	if h1 < 6 {
		h1 = h1 + 12
	} else {
		h1 = h1
	}
	if h2 < 6 {
		h2 = h2 + 12
	} else {
		h2 = h2
	}
	totalMasuk := (h1 * 60) + m1
	totalKeluar := (h2 * 60) + m2
	durasiTotal := totalKeluar - totalMasuk

	durasiJam := durasiTotal / 60
	durasiMenit := durasiTotal % 60

	fmt.Printf("Parking duration is %d hour %d minutes\n", durasiJam, durasiMenit)
}
