package main

import "fmt"

func main() {
	var input int
	var max, min int
	var first bool = true
	for {
		fmt.Scan(&input)

		if input == 0 {
			break
		}

		if first {
			max = input
			min = input
			first = false
		} else {
			if input > max {
				max = input
			}
			if input < min {
				min = input
			}
		}
	}

	if !first {
		fmt.Printf("%d %d\n", max, min)
	} else {
		fmt.Println("Tidak ada data")
	}
}
