package main

import "fmt"

func main() {
	var jumlahLangkah int
	var panjangLangkahCm int
	var kaloriPerKm float64
	var totalJarak float64
	var totalJarakCm int
	var totalKalori float64
	fmt.Print("Masukan jumlah langkah,panjang langkah dan kalori perkm: ")
	fmt.Scanln(&jumlahLangkah, &panjangLangkahCm, &kaloriPerKm)
	//ubah cm ke km
	totalJarakCm = jumlahLangkah * panjangLangkahCm

	totalJarak = float64(totalJarakCm) / 100000.0
	totalKalori = totalJarak * kaloriPerKm

	fmt.Printf("Distance = %.2f km, Calories = %.2f kcal\n", totalJarak, totalKalori)

}
