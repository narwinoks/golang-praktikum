package main

import "fmt"

func main() {
	var jenis string
	var jumlah, totalSD, totalSMP, totalSMA int

	for {
		fmt.Scan(&jenis, &jumlah)
		if jenis == "x" && jumlah == 0 {
			break
		}

		if jenis == "sd" {
			totalSD += jumlah
		} else if jenis == "smp" {
			totalSMP += jumlah
		} else if jenis == "sma" {
			totalSMA += jumlah
		}
	}
	fmt.Printf("%d %d %d\n", totalSD, totalSMP, totalSMA)
}
