package main

import "fmt"

func main() {
	var bilangan int = 2
	var jumlah int = 0
	for {

		jumlah = jumlah + bilangan
		bilangan++
		if bilangan > 50 {
			break
		}
	}
	fmt.Println(jumlah)
}
