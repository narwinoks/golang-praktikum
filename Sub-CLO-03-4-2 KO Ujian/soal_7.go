package main

import (
	"fmt"
)

func main() {
	var input int
	var positif, negatif int
	for {
		fmt.Scan(&input)

		if input == 0 {
			break
		}

		if input > 0 {
			positif++
		} else if input < 0 {
			negatif++
		}
	}

	fmt.Printf("%d %d\n", positif, negatif)
}
