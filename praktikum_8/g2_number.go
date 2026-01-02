package main

import "fmt"

func main() {
	var text string

	fmt.Scan(&text)

	count := 0

	for i := 0; i < len(text); i++ {

		char := text[i]

		switch {
		case char >= '0' && char <= '9':
			count++
		}
	}

	fmt.Printf("%d digit numbers\n", count)
}
