package main

import "fmt"

func main() {
	var jumlah int = 0
	var bilangan int = 2

	for {
		jumlah = jumlah + bilangan
		bilangan++

		if bilangan > 50 {
			break
		}
	}

	fmt.Println(jumlah)
}
