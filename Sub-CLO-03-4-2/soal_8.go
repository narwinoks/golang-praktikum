package main

import (
	"fmt"
)

func main() {
	var jenis string
	var jumlah, totalBagus, totalCacat int

	for {
		fmt.Scan(&jenis, &jumlah)

		if jenis == "selesai" && jumlah == 0 {
			break
		}

		if jenis == "bagus" {
			totalBagus += jumlah
		} else if jenis == "cacat" {
			totalCacat += jumlah
		}
	}

	fmt.Printf("%d %d\n", totalBagus, totalCacat)
}
