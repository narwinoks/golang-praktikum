package main

import "fmt"

func main() {
	var bilanganSekarang int
	var bilanganSebelumnya int
	var hitung int = 0

	for i := 0; i < 10; i++ {

		fmt.Scanln(&bilanganSekarang)

		if i > 0 {
			if bilanganSekarang > bilanganSebelumnya {
				break
			}
		}

		hitung++

		bilanganSebelumnya = bilanganSekarang
	}

	fmt.Println(hitung)
}
