package main

import "fmt"

func main() {
	var input1, input2 int
	fmt.Scan(&input1)
	fmt.Scan(&input2)
	for input1 <= input2 {
		fmt.Print(input1)
		if input1 < input2 {
			fmt.Print(" ")
		}
		input1++
	}
}
