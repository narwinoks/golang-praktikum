package main

import "fmt"

func main() {
	var totalBelanja int
	var isNational bool
	var jarak int
	var isCash bool
	fmt.Scan(&totalBelanja, &isNational, &jarak, &isCash)
	if jarak >= 10 {
		fmt.Println("Cashback IDR 10.000")
	}
	if isNational {
		fmt.Println("Discount IDR 15.000")
	}
	if totalBelanja >= 100000 {
		fmt.Println("Cashback 25%")
	}
	if !isCash {
		fmt.Println("Discount 50%")
	}
}
