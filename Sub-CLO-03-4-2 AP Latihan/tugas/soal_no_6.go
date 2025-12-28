package main

import (
	"fmt"
)

func main() {
	var node string
	fmt.Scan(&node)

	if node == "A" {
		fmt.Println("akar")
	} else if node == "B" || node == "E" {
		fmt.Println("verteks dalam")
	} else if node == "C" || node == "D" || node == "F" || node == "G" {
		fmt.Println("daun")
	}
}
