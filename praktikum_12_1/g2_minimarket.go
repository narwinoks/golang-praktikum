package main

import "fmt"

func main() {
	var n int
	var targetStore string

	fmt.Scan(&n, &targetStore)

	var currentStore string
	var count int = 0
	var i int = 0

	for i < n {
		fmt.Scan(&currentStore)

		if currentStore == targetStore {
			count++
		}

		i++
	}

	fmt.Println(targetStore, count)
}
