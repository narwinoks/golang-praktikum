package main

import (
	"fmt"
)

func main() {
	var bilangan int
	var genap, ganjil int

	for i := 0; i < 4; i++ {
		fmt.Scan(&bilangan)

		if bilangan%2 == 0 {
			genap++
		} else {
			ganjil++
		}
	}

	fmt.Printf("%d %d\n", genap, ganjil)
}
