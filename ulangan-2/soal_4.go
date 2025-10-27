package main

import "fmt"

func main() {
	var jumlah int = 0
	for bilangan := 0; bilangan <= 99; bilangan++ {
		if bilangan%2 != 0 {
			jumlah = jumlah + bilangan

		}
	}
	fmt.Println(jumlah)
}
