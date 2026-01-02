package main

import "fmt"

func main() {
	var token int
	var found bool = false
	for !found {
		fmt.Scan(&token)

		digit2 := (token / 100) % 10

		digit4 := token % 10

		sum := digit2 + digit4

		fmt.Printf("%d ", sum)

		if sum == 10 {
			found = true
		}
	}
}
