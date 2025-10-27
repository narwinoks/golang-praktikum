package main

import "fmt"

func main() {
	var jumlah int = 0
	var i int = 1

	for {

		jumlah += i
		i++
		if i > 100 {
			break
		}
	}

	fmt.Println(jumlah)
}
