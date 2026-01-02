package main

import "fmt"

func main() {
	var n int
	var grade string
	var sks int
	var totalSKS int = 0
	var totalMutu float64 = 0.0
	var gpa float64
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&grade, &sks)

		var bobot float64
		var isLulus bool = true

		if grade == "A" {
			bobot = 4.0
		} else if grade == "AB" {
			bobot = 3.5
		} else if grade == "B" {
			bobot = 3.0
		} else if grade == "BC" {
			bobot = 2.5
		} else if grade == "C" {
			bobot = 2.0
		} else if grade == "D" {
			bobot = 1.0
		} else {
			bobot = 0.0
			isLulus = false
		}

		if isLulus {
			totalMutu = totalMutu + (bobot * float64(sks))
			totalSKS = totalSKS + sks
		} else {
		}
	}

	if totalSKS > 0 {
		gpa = totalMutu / float64(totalSKS)
	} else {
		gpa = 0.00
	}

	fmt.Printf("The GPA is %.2f with %d completed credits\n", gpa, totalSKS)
}
