package main

import "fmt"

func main() {
	var jam, menit, biaya int
	fmt.Scan(&jam, &menit)
	if menit > 0 {
		jam = jam + 1
	} else {
		jam = jam
	}

	if jam <= 1 {
		biaya = 5000
	} else {
		biaya = 5000 + ((jam - 1) * 3000)
	}
	fmt.Printf("Parking fee is IDR %d\n", biaya)
}
