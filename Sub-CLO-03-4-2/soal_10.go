package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var total int = 0

	fmt.Scan(&input)

	for _, char := range input {
		strChar := string(char)

		if strChar == "." {
			break
		}

		angka, err := strconv.Atoi(strChar)
		if err == nil {
			total += angka
		}
	}

	fmt.Println(total)
}
