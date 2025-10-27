package main

import "fmt"

func main() {
	var bilangan int = 1
	var jumlah int = 0

	for {
		jumlah = jumlah + bilangan

		bilangan = bilangan + 2

		if bilangan > 99 {
			break
		}
	}

	fmt.Println(jumlah)
}
