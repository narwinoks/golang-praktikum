package main

import (
	"fmt"
)

func main() {
	var n int
	var daduA, daduB, daduC int
	var totalA, totalB, totalC int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&daduA, &daduB, &daduC)

		totalA += daduA
		totalB += daduB
		totalC += daduC
	}

	if totalA > totalB && totalA > totalC {
		fmt.Printf("A %d\n", totalA)
	} else if totalB > totalA && totalB > totalC {
		fmt.Printf("B %d\n", totalB)
	} else {
		fmt.Printf("C %d\n", totalC)
	}
}
