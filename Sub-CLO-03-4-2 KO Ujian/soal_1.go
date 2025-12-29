package main

import "fmt"

func main() {
	var n, suaraA, suaraB int

	fmt.Scan(&n, &suaraA, &suaraB)

	totalSuara := suaraA + suaraB

	syaratMinimal := float64(n) * 0.6

	if float64(totalSuara) >= syaratMinimal {
		if suaraA > suaraB {
			fmt.Println("Kandidat A menang")
		} else if suaraB > suaraA {
			fmt.Println("Kandidat B menang")
		} else {
			fmt.Println("Tidak ada pemenang")
		}
	} else {
		fmt.Println("Tidak ada pemenang")
	}
}
