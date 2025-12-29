package main

import (
	"fmt"
)

func main() {
	var d1, d2 int
	var count int = 0

	for {
		_, err := fmt.Scan(&d1, &d2)

		if err != nil {
			break
		}

		if d1%2 != 0 && d2%2 != 0 {
			break
		}

		if d1%2 == 0 && d2%2 == 0 {
			count++
		}
	}

	fmt.Println(count)
}
