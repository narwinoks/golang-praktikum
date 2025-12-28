package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var total int = 0
	var selesai bool = false

	for !selesai {
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}

		for _, char := range input {
			karakterString := string(char)

			if karakterString == "." {
				selesai = true
				break
			}

			angka, err := strconv.Atoi(karakterString)

			if err == nil {
				total += angka
			}
		}
	}

	fmt.Println(total)
}
