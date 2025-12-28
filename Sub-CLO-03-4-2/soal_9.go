package main

import "fmt"

func main() {
	var n, i int
	var tipeKendaraan string

	var countA, countB, countC int
	var totalPendapatan int = 0

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&tipeKendaraan)

		if tipeKendaraan == "A" {
			countA++
			totalPendapatan += 10000
		} else if tipeKendaraan == "B" {
			countB++
			totalPendapatan += 23000
		} else if tipeKendaraan == "C" {
			countC++
			totalPendapatan += 45000
		}
	}

	fmt.Printf("%d %d %d %d\n", totalPendapatan, countA, countB, countC)
}
