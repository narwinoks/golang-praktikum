package main

import "fmt"

func main() {
	var golKita, golLawan int

	// Membaca input (baris 1: gol kita, baris 2: gol lawan)
	fmt.Scan(&golKita, &golLawan)

	if golKita > golLawan {
		fmt.Println("menang")
	} else if golKita < golLawan {
		fmt.Println("kalah")
	} else {
		fmt.Println("draw")
	}
}
