package main

import "fmt"

func main() {
	var input int
	var total int = 0

	for {
		_, err := fmt.Scan(&input)

		if err != nil {
			break
		}

		if input < 0 {
			break
		}

		if input%2 != 0 && input%3 == 0 {
			total += input
		}
	}
	fmt.Println(total)
}
